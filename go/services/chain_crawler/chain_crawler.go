package chain_crawler

import (
	"encoding/json"
	"errors"
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
const maxErrorCntUntilNotification = 96

type ChainCrawler struct {
	httpClient            *http.Client
	chainManager          *database.ChainManager
	chainProposalManager  *database.ChainProposalManager
	addressTrackerManager *database.AddressTrackerManager
	notifier              notifier.ChainNotifier
	assetsPath            string
	errorCnt              map[int]int
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
		errorCnt:              make(map[int]int),
	}
}

func (c *ChainCrawler) getJson(url string, target interface{}) (int, error) {
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return resp.StatusCode, errors.New(resp.Status)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	return resp.StatusCode, json.NewDecoder(resp.Body).Decode(target)
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

func (c *ChainCrawler) reportErrorIfNecessary(chainEnt *ent.Chain, url string, err error) {
	c.errorCnt[chainEnt.ID]++
	if c.errorCnt[chainEnt.ID]%maxErrorCntUntilNotification == 0 { // report every `maxErrorCntUntilNotification` times
		log.Sugar.Errorf("Error calling `%v`: %v", url, err)
	}
}

func (c *ChainCrawler) resetErrorCount(chainEnt *ent.Chain) {
	c.errorCnt[chainEnt.ID] = 0
}

func (c *ChainCrawler) addProposals(chainEnt *ent.Chain, url string) []ProposalInfo {
	var resp types.ChainProposalsResponse
	_, err := c.getJson(url, &resp)
	if err != nil {
		c.reportErrorIfNecessary(chainEnt, url, err)
		return nil
	} else {
		c.resetErrorCount(chainEnt)
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
	_, err := c.getJson(url, &resp)
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

func (c *ChainCrawler) AddOrUpdateChains() {
	log.Sugar.Debug("Updating chains")
	var chainInfo types.ChainInfo
	_, err := c.getJson("https://chains.cosmos.directory/", &chainInfo)
	if err != nil {
		log.Sugar.Errorf("Error calling reg.ListChains: %v", err)
	}

	for _, chain := range chainInfo.Chains {
		var found = false
		for _, entChain := range c.chainManager.All() {
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
			entChain := c.chainManager.Create(&chain, thumbnailUrl)
			url := fmt.Sprintf(urlProposals, getChainPath(entChain))
			c.addProposals(entChain, url)
		}
	}
}

func getChainPath(chainEnt *ent.Chain) string {
	chainPath := chainEnt.Path
	if chainPath == "" {
		chainPath = chainEnt.Name
	}
	return chainPath
}

func (c *ChainCrawler) UpdateProposals() {
	log.Sugar.Debug("Updating chain proposals")
	for _, chainEnt := range c.chainManager.All() {
		log.Sugar.Debugf("Updating proposals for chain %v", chainEnt.PrettyName)
		for _, entProposal := range c.chainProposalManager.VotingPeriodExpired(chainEnt) {
			url := fmt.Sprintf(urlProposals+"/%v", getChainPath(chainEnt), entProposal.ProposalID)
			c.updateProposal(chainEnt, url)
		}

		url := fmt.Sprintf(urlProposals+"?proposal_status=%v", getChainPath(chainEnt), cosmossdktypes.StatusVotingPeriod)
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
	for _, data := range c.addressTrackerManager.GetAllUnnotifiedTrackers() {
		url := fmt.Sprintf(urlVote, getChainPath(data.ChainProposal.Edges.Chain), data.ChainProposal.ProposalID, data.AddressTracker.Address)
		var voteResponse types.ChainProposalVoteResponse
		statusCode, err := c.getJson(url, &voteResponse)
		if err != nil && statusCode == 400 {
			c.notifier.SendVoteReminder(data.ChainProposal.Edges.Chain, data.ChainProposal)
		} else if err != nil {
			c.reportErrorIfNecessary(data.ChainProposal.Edges.Chain, url, err)
			continue
		} else {
			if voteResponse.Vote.Option.ToCosmosType() == cosmossdktypes.OptionEmpty {
				c.notifier.SendVoteReminder(data.ChainProposal.Edges.Chain, data.ChainProposal)
			}
		}
		c.resetErrorCount(data.ChainProposal.Edges.Chain)
		c.addressTrackerManager.SetNotified(data)
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
