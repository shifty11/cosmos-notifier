package chain_crawler

import (
	"encoding/json"
	"fmt"
	"github.com/shifty11/dao-dao-notifier/common"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/types"
	"net/http"
	"time"
)

type ChainCrawler struct {
	client               *http.Client
	chainManager         *database.ChainManager
	chainProposalManager *database.ChainProposalManager
	assetsPath           string
}

func NewChainCrawler(dbManagers *database.DbManagers, assetsPath string) *ChainCrawler {
	var client = &http.Client{Timeout: 10 * time.Second}
	return &ChainCrawler{
		client:               client,
		chainManager:         dbManagers.ChainManager,
		chainProposalManager: dbManagers.ChainProposalManager,
		assetsPath:           assetsPath,
	}
}

func (cc *ChainCrawler) getJson(url string, target interface{}) error {
	resp, err := cc.client.Get(url)
	if err != nil {
		return nil
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

func (cc *ChainCrawler) addProposals(entChain *ent.Chain) {
	var resp types.ChainProposalsResponse
	url := fmt.Sprintf("https://rest.cosmos.directory/%v/cosmos/gov/v1beta1/proposals", entChain.Name)
	err := cc.getJson(url, &resp)
	if err != nil {
		log.Sugar.Errorf("Error calling `%v`: %v", url, err)
		return
	}

	for _, proposal := range resp.Proposals {
		cc.chainProposalManager.CreateOrUpdate(entChain, &proposal)
	}
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
			cc.addProposals(entChain)
		}
	}
}
