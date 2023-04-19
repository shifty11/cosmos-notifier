package chain_crawler

import (
	"fmt"
	cosmossdktypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/robfig/cron/v3"
	"github.com/shifty11/cosmos-notifier/common"
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/log"
	"github.com/shifty11/cosmos-notifier/notifier"
	"github.com/shifty11/cosmos-notifier/types"
	"net/http"
	"time"
)

const urlProposals = "https://rest.cosmos.directory/%v/cosmos/gov/v1beta1/proposals"
const urlVote = urlProposals + "/%v/votes/%v"

type ChainCrawler struct {
	httpClient            *http.Client
	chainManager          *database.ChainManager
	chainProposalManager  *database.ChainProposalManager
	addressTrackerManager *database.AddressTrackerManager
	notifier              notifier.ChainNotifier
	assetsPath            string
	errorReporter         common.ErrorReporter
}

func NewChainCrawler(dbManagers *database.DbManagers, notifier notifier.ChainNotifier, assetsPath string) *ChainCrawler {
	var client = &http.Client{Timeout: 10 * time.Second}
	return &ChainCrawler{
		httpClient:            client,
		chainManager:          dbManagers.ChainManager,
		chainProposalManager:  dbManagers.ChainProposalManager,
		addressTrackerManager: dbManagers.AddressTrackerManager,
		notifier:              notifier,
		assetsPath:            assetsPath,
		errorReporter:         common.NewErrorReporter(common.DefaultMaxErrorCntUntilReport),
	}
}

func (c *ChainCrawler) imageManager(chain *types.Chain) *common.ImageManager {
	return common.NewImageManager(
		chain.Name,
		chain.PrettyName,
		chain.Image,
		c.assetsPath,
		"shared/chains/",
		100,
		100,
	)
}

func (c *ChainCrawler) downloadThumbnail(chain *types.Chain) string {
	im := c.imageManager(chain)
	err := im.DownloadAndCreateThumbnail()
	if err != nil {
		log.Sugar.Infof("while downloading image for chain %v: %v", chain.PrettyName, err)
	} else {
		return im.ThumbnailUrl
	}
	return ""
}

func (c *ChainCrawler) doesThumbnailExist(chain *types.Chain) bool {
	return c.imageManager(chain).DoesExist()
}

type ProposalInfo struct {
	proposal *ent.ChainProposal
	status   database.ProposalStatus
}

func (c *ChainCrawler) addProposals(chainEnt *ent.Chain, url string) []ProposalInfo {
	var resp types.ChainProposalsResponse
	_, err := common.GetJson(c.httpClient, url, &resp)
	if err != nil {
		c.errorReporter.ReportErrorIfNecessary(
			chainEnt.ID,
			fmt.Sprintf("Error calling `%v`: %v", url, err),
		)
		return nil
	} else {
		c.errorReporter.ResetErrorCount(chainEnt.ID)
	}

	var props []ProposalInfo
	for _, proposal := range resp.Proposals {
		prop, status := c.chainProposalManager.CreateOrUpdate(chainEnt, &proposal)
		props = append(props, ProposalInfo{
			proposal: prop,
			status:   status,
		})
	}
	return props
}

func (c *ChainCrawler) updateProposal(entChain *ent.Chain, url string) {
	var resp types.ChainProposalResponse
	_, err := common.GetJson(c.httpClient, url, &resp)
	if err != nil {
		log.Sugar.Errorf("Error calling `%v`: %v", url, err)
		return
	}
	c.chainProposalManager.CreateOrUpdate(entChain, &resp.Proposal)
}

func (c *ChainCrawler) chainNeedsUpdate(entChain *ent.Chain, chainInfo *types.Chain) bool {
	return entChain.ChainID != chainInfo.ChainId ||
		entChain.Name != chainInfo.Name ||
		entChain.PrettyName != chainInfo.PrettyName ||
		entChain.Path != chainInfo.Path ||
		entChain.Display != chainInfo.Display ||
		entChain.Bech32Prefix != chainInfo.Bech32Prefix ||
		entChain.ImageURL != chainInfo.Image ||
		(chainInfo.Image != "" && entChain.ThumbnailURL == "") ||
		!c.doesThumbnailExist(chainInfo)
}

func (c *ChainCrawler) isChainValid(chainInfo *types.Chain) bool {
	return chainInfo.NetworkType == "mainnet" && chainInfo.ChainId != "" && chainInfo.Path != "" && chainInfo.Bech32Prefix != ""
}

func (c *ChainCrawler) AddOrUpdateChains() {
	log.Sugar.Debug("Updating chains")
	var chainInfo types.ChainInfo
	_, err := common.GetJson(c.httpClient, "https://chains.cosmos.directory/", &chainInfo)
	if err != nil {
		log.Sugar.Errorf("Error calling reg.ListChains: %v", err)
	}

	for _, chain := range chainInfo.Chains {
		if !c.isChainValid(&chain) {
			continue
		}
		var found = false
		for _, entChain := range c.chainManager.QueryAll() {
			if entChain.Name == chain.Name {
				found = true
				if c.chainNeedsUpdate(entChain, &chain) {
					thumbnailUrl := entChain.ThumbnailURL
					if !c.doesThumbnailExist(&chain) {
						thumbnailUrl = c.downloadThumbnail(&chain)
					}
					c.chainManager.Update(entChain, &chain, thumbnailUrl)
				}
				break
			}
		}
		if !found && chain.NetworkType == "mainnet" {
			thumbnailUrl := c.downloadThumbnail(&chain)
			chainEnt := c.chainManager.Create(&chain, thumbnailUrl)
			url := fmt.Sprintf(urlProposals, chainEnt.Path)
			c.addProposals(chainEnt, url)
		}
	}
}

func (c *ChainCrawler) UpdateProposals() {
	log.Sugar.Debug("Updating chain proposals")
	for _, chainEnt := range c.chainManager.QueryEnabled() {
		log.Sugar.Debugf("Updating proposals for chain %v", chainEnt.PrettyName)
		for _, entProposal := range c.chainProposalManager.QueryVotingPeriodExpired(chainEnt) {
			url := fmt.Sprintf(urlProposals+"/%v", chainEnt.Path, entProposal.ProposalID)
			c.updateProposal(chainEnt, url)
		}

		url := fmt.Sprintf(urlProposals+"?proposal_status=%v", chainEnt.Path, cosmossdktypes.StatusVotingPeriod)
		props := c.addProposals(chainEnt, url)
		for _, prop := range props {
			if chainEnt.IsEnabled && prop.status == database.ProposalCreated {
				c.notifier.Notify(chainEnt, prop.proposal)
			}
		}
	}
}

func (c *ChainCrawler) CheckForVotingReminders() {
	log.Sugar.Info("Checking for voting reminders")
	for _, data := range c.addressTrackerManager.QueryUnnotifiedTrackers() {
		url := fmt.Sprintf(urlVote, data.ChainProposal.Edges.Chain.Path, data.ChainProposal.ProposalID, data.AddressTracker.Address)
		var voteResponse types.ChainProposalVoteResponse
		statusCode, err := common.GetJson(c.httpClient, url, &voteResponse)
		if err != nil && statusCode == 400 {
			c.notifier.SendVoteReminder(data)
		} else if err != nil {
			c.errorReporter.ReportErrorIfNecessary(
				data.ChainProposal.Edges.Chain.ID,
				fmt.Sprintf("Error calling `%v`: %v", url, err),
			)
			continue
		} else {
			if voteResponse.Vote.Option.ToCosmosType() == cosmossdktypes.OptionEmpty {
				c.notifier.SendVoteReminder(data)
			}
		}
		c.errorReporter.ResetErrorCount(data.ChainProposal.Edges.Chain.ID)
		c.addressTrackerManager.UpdateSetNotified(data)
	}
}

func (c *ChainCrawler) ScheduleCrawl() {
	log.Sugar.Info("Scheduling chain crawl")
	cr := cron.New()
	_, err := cr.AddFunc("*/15 * * * *", func() { c.UpdateProposals() }) // every 15min
	if err != nil {
		log.Sugar.Errorf("while executing 'UpdateProposals' via cron: %v", err)
	}
	_, err = cr.AddFunc("5-50/15 * * * *", func() { c.CheckForVotingReminders() }) // every 15min but not at the same time as UpdateProposals
	if err != nil {
		log.Sugar.Errorf("while executing 'CheckForVotingReminders' via cron: %v", err)
	}
	_, err = cr.AddFunc("0 9 * * *", func() { c.AddOrUpdateChains() }) // every day at 9:00
	if err != nil {
		log.Sugar.Errorf("while executing 'AddOrUpdateChains' via cron: %v", err)
	}
	cr.Start()
}
