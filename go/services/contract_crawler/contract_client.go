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
	"strings"
)

type ContractClient struct {
	URL             string
	ContractAddress string
	RpcEndpoint     string
}

func NewContractClient(URL string, contractAddress string, rpcEndpoint string) *ContractClient {
	if strings.HasSuffix(URL, "/") {
		URL = URL[:len(URL)-1]
	}
	return &ContractClient{
		URL:             URL,
		ContractAddress: contractAddress,
		RpcEndpoint:     rpcEndpoint,
	}
}

func (cc *ContractClient) configV2() (*types.ContractData, error) {
	resp, err := cc.querySmartContract("{\"get_config\":{}}")
	if err != nil {
		return nil, err
	}

	var config types.Config
	err = json.Unmarshal(resp, &config)
	if err != nil {
		return nil, err
	}
	return config.ToContractData(cc.ContractAddress, cc.RpcEndpoint), nil
}

func (cc *ContractClient) configV1() (*types.ContractData, error) {
	resp, err := cc.querySmartContract("{\"config\":{}}")
	if err != nil {
		return nil, err
	}
	var config types.ConfigV1
	err = json.Unmarshal(resp, &config)
	if err != nil {
		return nil, err
	}
	return config.ToContractData(cc.ContractAddress, cc.RpcEndpoint), nil
}

func (cc *ContractClient) config(contractVersion types.ContractVersion) (*types.ContractData, error) {
	switch contractVersion {
	case types.ContractVersionV1:
		return cc.configV1()
	case types.ContractVersionV2:
		return cc.configV2()
	default:
		result, err := cc.configV2()
		if err != nil {
			if errors.Is(err, UnknownVariantError) {
				return cc.configV1()
			}
			return nil, err
		}
		return result, nil
	}
}

func (cc *ContractClient) proposals(query string) (*types.ProposalList, error) {
	if query == "" {
		query = "{\"list_proposals\":{}}"
	}
	resp, err := cc.querySmartContract(query)
	if err != nil {
		if errors.Is(err, UnknownVariantError) {
			proposalModules, _ := cc.proposalModules()
			if len(proposalModules) > 0 {
				return NewContractClient(cc.URL, proposalModules[0], cc.RpcEndpoint).proposals("")
			}
			return nil, errors.New(fmt.Sprintf("found no proposal_modules for %v", cc.ContractAddress))
		}
		return nil, err
	}

	var proposals types.ProposalList
	err = json.Unmarshal(resp, &proposals)
	if err != nil {
		return nil, err
	}
	if len(proposals.Proposals) > 0 && proposals.Proposals[0].Status == "" {
		var propsV1 types.ProposalListV1
		err = json.Unmarshal(resp, &propsV1)
		if err != nil {
			log.Sugar.Fatal(err)
		}

		var propsList = propsV1.ToProposalList()
		if len(propsList.Proposals) == 0 || propsList.Proposals[0].Status == "" {
			return nil, errors.New(fmt.Sprintf("could not decode proposals of contract %v", cc.ContractAddress))
		}
		return propsList, nil
	}
	return &proposals, nil
}

func (cc *ContractClient) proposalModules() ([]string, error) {
	resp, err := cc.querySmartContract("{\"proposal_modules\":{}}")
	if err != nil {
		return nil, err
	}
	var proposalModules []string
	err = json.Unmarshal(resp, &proposalModules)
	return proposalModules, err
}

func (cc *ContractClient) proposal(proposalId int) (*types.Proposal, error) {
	resp, err := cc.querySmartContract(fmt.Sprintf("{\"proposal\":{\"proposal_id\": %v}}", proposalId))
	if err != nil {
		return nil, err
	}
	var proposal types.Proposal
	err = json.Unmarshal(resp, &proposal)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	return &proposal, nil
}

var UnknownVariantError = errors.New("unknown variant")

func (cc *ContractClient) querySmartContract(query string) ([]byte, error) {
	values := map[string]string{"contractAddress": cc.ContractAddress, "rpcEndpoint": cc.RpcEndpoint, "query": query}
	jsonData, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%v/query_smart_contract", cc.URL), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotAcceptable {
			return nil, UnknownVariantError
		}
		respErr, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error read response from /query_smart_contract: %v %v", resp.StatusCode, err))
		}
		return nil, errors.New(fmt.Sprintf("error querying /query_smart_contract: %v %v", resp.StatusCode, respErr))
	}

	return io.ReadAll(resp.Body)
}
