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

type ChainCrawler struct {
	client               *http.Client
	chainManager         *database.ChainManager
	chainProposalManager *database.ChainProposalManager
	notifier             *notifier.ChainNotifier
	assetsPath           string
}

func NewChainCrawler(dbManagers *database.DbManagers, notifier *notifier.ChainNotifier, assetsPath string) *ChainCrawler {
	var client = &http.Client{Timeout: 10 * time.Second}
	return &ChainCrawler{
		client:               client,
		chainManager:         dbManagers.ChainManager,
		chainProposalManager: dbManagers.ChainProposalManager,
		notifier:             notifier,
		assetsPath:           assetsPath,
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

func (c *ChainCrawler) downloadImage(chain *types.Chain) string {
	im := common.NewImageManager(
		chain.Name,
		chain.PrettyName,
		chain.Image,
		c.assetsPath,
		"images/chains/",
		100,
		100,
	)
	err := im.DownloadAndCreateThumbnail()
	if err != nil {
		log.Sugar.Infof("while downloading image for chain %v: %v", chain.PrettyName, err)
	} else {
		return im.ThumbnailUrl
	}
	return ""
}

func (c *ChainCrawler) addProposals(entChain *ent.Chain, url string) []*ent.ChainProposal {
	var resp types.ChainProposalsResponse
	err := c.getJson(url, &resp)
	if err != nil {
		log.Sugar.Errorf("Error calling `%v`: %v", url, err)
		return nil
	}

	var props []*ent.ChainProposal
	for _, proposal := range resp.Proposals {
		prop, _ := c.chainProposalManager.CreateOrUpdate(entChain, &proposal)
		props = append(props, prop)
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

func (c *ChainCrawler) AddOrUpdateChains() {
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
				if (entChain.Name != chain.Name) ||
					(entChain.PrettyName != chain.PrettyName) ||
					(entChain.ImageURL != chain.Image) ||
					(chain.Image != "" && entChain.ThumbnailURL == "") {
					thumbnailUrl := c.downloadImage(&chain)
					c.chainManager.Update(entChain, &chain, thumbnailUrl)
				}
				thumbnailUrl := c.downloadImage(&chain)
				c.chainManager.Update(entChain, &chain, thumbnailUrl)
				break
			}
		}
		if !found && chain.NetworkType == "mainnet" {
			thumbnailUrl := c.downloadImage(&chain)
			entChain := c.chainManager.Create(&chain, thumbnailUrl)
			url := fmt.Sprintf(urlProposals, entChain.Name)
			c.addProposals(entChain, url)
		}
	}
}

func (c *ChainCrawler) UpdateProposals() {
	for _, entChain := range c.chainManager.All() {
		for _, entProposal := range c.chainProposalManager.InVotingPeriod(entChain) {
			url := fmt.Sprintf(urlProposals+"/%v", entChain.Name, entProposal.ProposalID)
			c.updateProposal(entChain, url)
		}

		url := fmt.Sprintf(urlProposals+"?proposal_status=%v", entChain.Name, cosmossdktypes.StatusVotingPeriod)
		props := c.addProposals(entChain, url)
		for _, prop := range props {
			if entChain.IsEnabled {
				c.notifier.Notify(entChain, prop)
			}
		}
	}
}

func (c *ChainCrawler) ScheduleCrawl() {
	log.Sugar.Info("Scheduling chain crawl")
	cr := cron.New()
	_, err := cr.AddFunc("@every 15min", func() { c.UpdateProposals() })
	if err != nil {
		log.Sugar.Errorf("while executing 'updateContracts' via cron: %v", err)
	}
	cr.Start()
}