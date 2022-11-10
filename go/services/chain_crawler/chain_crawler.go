package chain_crawler

import (
	"encoding/json"
	"fmt"
	cosmossdktypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/robfig/cron/v3"
	"github.com/shifty11/dao-dao-notifier/common"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/notifier"
	"github.com/shifty11/dao-dao-notifier/types"

	"net/http"
	"time"
)

const urlProposals = "https://rest.cosmos.directory/%v/cosmos/gov/v1beta1/proposals"
const maxErrorCntUntilNotification = 96

type ChainCrawler struct {
	client               *http.Client
	chainManager         *database.ChainManager
	chainProposalManager *database.ChainProposalManager
	notifier             *notifier.ChainNotifier
	assetsPath           string
	errorCnt             map[int]int
}

func NewChainCrawler(dbManagers *database.DbManagers, notifier *notifier.ChainNotifier, assetsPath string) *ChainCrawler {
	var client = &http.Client{Timeout: 10 * time.Second}
	return &ChainCrawler{
		client:               client,
		chainManager:         dbManagers.ChainManager,
		chainProposalManager: dbManagers.ChainProposalManager,
		notifier:             notifier,
		assetsPath:           assetsPath,
		errorCnt:             make(map[int]int),
	}
}

func (c *ChainCrawler) getJson(url string, target interface{}) error {
	resp, err := c.client.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
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

func (c *ChainCrawler) addProposals(entChain *ent.Chain, url string) []ProposalInfo {
	var resp types.ChainProposalsResponse
	err := c.getJson(url, &resp)
	if err != nil {
		c.errorCnt[entChain.ID]++
		if c.errorCnt[entChain.ID]%maxErrorCntUntilNotification == 0 { // report every `maxErrorCntUntilNotification` times
			log.Sugar.Errorf("Error calling `%v`: %v", url, err)
		}
		return nil
	}

	var props []ProposalInfo
	for _, proposal := range resp.Proposals {
		prop, status := c.chainProposalManager.CreateOrUpdate(entChain, &proposal)
		props = append(props, ProposalInfo{
			proposal: prop,
			status:   status,
		})
	}
	return props
}

func (c *ChainCrawler) updateProposal(entChain *ent.Chain, url string) {
	var resp types.ChainProposalResponse
	err := c.getJson(url, &resp)
	if err != nil {
		log.Sugar.Errorf("Error calling `%v`: %v", url, err)
		return
	}
	c.chainProposalManager.CreateOrUpdate(entChain, &resp.Proposal)
}

func (c *ChainCrawler) chainNeedsUpdate(entChain *ent.Chain, chainInfo *types.Chain) bool {
	return entChain.Name != chainInfo.Name ||
		entChain.PrettyName != chainInfo.PrettyName ||
		entChain.Path != chainInfo.Path ||
		entChain.Display != chainInfo.Display ||
		entChain.ImageURL != chainInfo.Image ||
		(chainInfo.Image != "" && entChain.ThumbnailURL == "") ||
		!c.doesThumbnailExist(chainInfo)
}

func (c *ChainCrawler) AddOrUpdateChains() {
	log.Sugar.Debug("Updating chains")
	var chainInfo types.ChainInfo
	err := c.getJson("https://chains.cosmos.directory/", &chainInfo)
	if err != nil {
		log.Sugar.Errorf("Error calling reg.ListChains: %v", err)
	}

	for _, chain := range chainInfo.Chains {
		var found = false
		for _, entChain := range c.chainManager.All() {
			if entChain.ChainID == chain.ChainId {
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
			url := fmt.Sprintf(urlProposals, entChain.Name)
			c.addProposals(entChain, url)
		}
	}
}

func (c *ChainCrawler) UpdateProposals() {
	log.Sugar.Debug("Updating chain proposals")
	for _, entChain := range c.chainManager.All() {
		log.Sugar.Debugf("Updating proposals for chain %v", entChain.PrettyName)
		for _, entProposal := range c.chainProposalManager.VotingPeriodExpired(entChain) {
			url := fmt.Sprintf(urlProposals+"/%v", entChain.Name, entProposal.ProposalID)
			c.updateProposal(entChain, url)
		}

		urlChain := entChain.Path
		if urlChain == "" {
			urlChain = entChain.Name
		}
		url := fmt.Sprintf(urlProposals+"?proposal_status=%v", urlChain, cosmossdktypes.StatusVotingPeriod)
		props := c.addProposals(entChain, url)
		for _, prop := range props {
			if entChain.IsEnabled && prop.status == database.ProposalCreated {
				c.notifier.Notify(entChain, prop.proposal)
			}
		}
	}
}

func (c *ChainCrawler) ScheduleCrawl() {
	log.Sugar.Info("Scheduling chain crawl")
	cr := cron.New()
	_, err := cr.AddFunc("*/15 * * * *", func() { c.UpdateProposals() }) // every 15min
	if err != nil {
		log.Sugar.Errorf("while executing 'updateContracts' via cron: %v", err)
	}
	_, err = cr.AddFunc("0 9 * * *", func() { c.AddOrUpdateChains() }) // every day at 9:00
	if err != nil {
		log.Sugar.Errorf("while executing 'updateContracts' via cron: %v", err)
	}
	cr.Start()
}
