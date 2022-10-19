package contract_crawler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/types"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type ContractClient struct {
	URL             string
	ContractAddress string
}

func NewContractClient(URL string, contractAddress string) *ContractClient {
	if strings.HasSuffix(URL, "/") {
		URL = URL[:len(URL)-1]
	}
	return &ContractClient{
		URL:             URL,
		ContractAddress: contractAddress,
	}
}

func (cc *ContractClient) config() (*types.ContractData, error) {
	values := map[string]string{"contractAddress": cc.ContractAddress}
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%v/get_config", cc.URL), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotAcceptable {
			return cc.configV1(jsonData)
		}
		log.Sugar.Errorf("error querying /get_config: %v", resp.StatusCode)
		return nil, errors.New("error querying /get_config")
	}

	var config types.Config
	err = json.NewDecoder(resp.Body).Decode(&config)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	return config.ToContractData(cc.ContractAddress), nil
}

func (cc *ContractClient) configV1(jsonData []byte) (*types.ContractData, error) {
	resp, err := http.Post(fmt.Sprintf("%v/config", cc.URL), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Sugar.Errorf("error querying /config: %v", resp.StatusCode)
		return nil, errors.New("error querying /config")
	}

	var configV1 types.ConfigV1
	err = json.NewDecoder(resp.Body).Decode(&configV1)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	return configV1.ToContractData(cc.ContractAddress), nil
}

func (cc *ContractClient) proposals() (*types.ProposalList, error) {
	values := map[string]string{"contractAddress": cc.ContractAddress}
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%v/list_proposals", cc.URL), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotAcceptable {
			proposalModules := cc.proposalModules(jsonData)
			if len(proposalModules) > 0 {
				return NewContractClient(cc.URL, proposalModules[0]).proposals()
			}
			return nil, errors.New(fmt.Sprintf("found no proposal_modules for %v", cc.ContractAddress))
		}
		log.Sugar.Errorf("error querying /list_proposals: %v", resp.StatusCode)
		return nil, errors.New("error querying /list_proposals")
	}

	var proposals types.ProposalList
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	err = json.Unmarshal(body, &proposals)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}
	// if status is empty it means that the format of proposals is old and needs to be converted
	if len(proposals.Proposals) > 0 && proposals.Proposals[0].Status == "" {
		var propsV1 types.ProposalListV1
		err = json.Unmarshal(body, &propsV1)
		if err != nil {
			log.Sugar.Fatal(err)
			return nil, err
		}

		var propsList = propsV1.ToProposalList()
		if len(propsList.Proposals) == 0 || propsList.Proposals[0].Status == "" {
			return nil, errors.New(fmt.Sprintf("could not decode proposals of contract %v", cc.ContractAddress))
		}
		return propsList, nil
	}

	return &proposals, nil
}

func (cc *ContractClient) proposalModules(jsonData []byte) []string {
	resp, err := http.Post(fmt.Sprintf("%v/proposal_modules", cc.URL), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Sugar.Fatal(err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		log.Sugar.Errorf("error querying /proposal_modules: %v", resp.StatusCode)
		return nil
	}

	var proposalModules []string
	err = json.NewDecoder(resp.Body).Decode(&proposalModules)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil
	}
	return proposalModules
}

func (cc *ContractClient) proposal(proposalId int) (*types.Proposal, error) {
	values := map[string]string{"contractAddress": cc.ContractAddress, "id": strconv.Itoa(proposalId)}
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%v/proposal", cc.URL), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Sugar.Errorf("error querying /proposal: %v", resp.StatusCode)
		return nil, errors.New("error querying /proposal")
	}

	var proposal types.Proposal
	err = json.NewDecoder(resp.Body).Decode(&proposal)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	return &proposal, nil
}
