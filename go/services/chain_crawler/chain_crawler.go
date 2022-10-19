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

func (cc *ChainCrawler) getJson(url string, target interface{}) error {
	resp, err := cc.client.Get(url)
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

func (cc *ChainCrawler) downloadImage(chain *types.Chain) string {
	im := common.NewImageManager(
		chain.Name,
		chain.PrettyName,
		chain.Image,
		cc.assetsPath,
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

func (cc *ChainCrawler) addProposals(entChain *ent.Chain, url string) []*ent.ChainProposal {
	var resp types.ChainProposalsResponse
	err := cc.getJson(url, &resp)
	if err != nil {
		log.Sugar.Errorf("Error calling `%v`: %v", url, err)
		return nil
	}

	var props []*ent.ChainProposal
	for _, proposal := range resp.Proposals {
		prop, _ := cc.chainProposalManager.CreateOrUpdate(entChain, &proposal)
		props = append(props, prop)
	}
	return props
}

func (cc *ChainCrawler) updateProposal(entChain *ent.Chain, url string) {
	var resp types.ChainProposalResponse
	err := cc.getJson(url, &resp)
	if err != nil {
		log.Sugar.Errorf("Error calling `%v`: %v", url, err)
		return
	}
	cc.chainProposalManager.CreateOrUpdate(entChain, &resp.Proposal)
}

func (cc *ChainCrawler) AddOrUpdateChains() {
	var chainInfo types.ChainInfo
	err := cc.getJson("https://chains.cosmos.directory/", &chainInfo)
	if err != nil {
		log.Sugar.Errorf("Error calling reg.ListChains: %v", err)
	}

	for _, chain := range chainInfo.Chains {
		var found = false
		for _, entChain := range cc.chainManager.All() {
			if entChain.ChainID == chain.ChainId {
				found = true
				if (entChain.Name != chain.Name) ||
					(entChain.PrettyName != chain.PrettyName) ||
					(entChain.ImageURL != chain.Image) ||
					(chain.Image != "" && entChain.ThumbnailURL == "") {
					thumbnailUrl := cc.downloadImage(&chain)
					cc.chainManager.Update(entChain, &chain, thumbnailUrl)
				}
				break
			}
		}
		if !found && chain.NetworkType == "mainnet" {
			thumbnailUrl := cc.downloadImage(&chain)
			entChain := cc.chainManager.Create(&chain, thumbnailUrl)
			url := fmt.Sprintf(urlProposals, entChain.Name)
			cc.addProposals(entChain, url)
		}
	}
}

func (cc *ChainCrawler) UpdateProposals() {
	for _, entChain := range cc.chainManager.All() {
		for _, entProposal := range cc.chainProposalManager.InVotingPeriod(entChain) {
			url := fmt.Sprintf(urlProposals+"/%v", entChain.Name, entProposal.ProposalID)
			cc.updateProposal(entChain, url)
		}

		url := fmt.Sprintf(urlProposals+"?proposal_status=%v", entChain.Name, cosmossdktypes.StatusVotingPeriod)
		props := cc.addProposals(entChain, url)
		for _, prop := range props {
			if entChain.IsEnabled {
				cc.notifier.Notify(entChain, prop)
			}
		}
	}
}

func (cc *ChainCrawler) ScheduleCrawl() {
	log.Sugar.Info("Scheduling chain crawl")
	cr := cron.New()
	_, err := cr.AddFunc("@every 15min", func() { cc.UpdateProposals() })
	if err != nil {
		log.Sugar.Errorf("while executing 'updateContracts' via cron: %v", err)
	}
	cr.Start()
}
