package service_crawler

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/types"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type ContractClient struct {
	ContractAddress string
}

func NewContractClient(contractAddress string) *ContractClient {
	return &ContractClient{
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

	resp, err := http.Post("http://localhost:8080/get_config", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotAcceptable {
			return cc.configV1(jsonData)
		}
		log.Sugar.Errorf("error querying /get_config (%v): %v", resp.StatusCode, resp.Body)
		return nil, errors.New("error querying /get_config")
	}

	var config types.Config
	err = json.NewDecoder(resp.Body).Decode(&config)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	return config.ToContractData(), nil
}

func (cc *ContractClient) configV1(jsonData []byte) (*types.ContractData, error) {
	resp, err := http.Post("http://localhost:8080/config", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Sugar.Errorf("error querying /config (%v): %v", resp.StatusCode, resp.Body)
		return nil, errors.New("error querying /config")
	}

	var configV1 types.ConfigV1
	err = json.NewDecoder(resp.Body).Decode(&configV1)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	return configV1.ToContractData(), nil
}

func (cc *ContractClient) proposals() (*types.ProposalList, error) {
	values := map[string]string{"contractAddress": cc.ContractAddress}
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	resp, err := http.Post("http://localhost:8080/list_proposals", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotAcceptable {
			proposalModules := cc.proposalModules(jsonData)
			if len(proposalModules) > 0 {
				return NewContractClient(proposalModules[0]).proposals()
			}
			return nil, errors.New(fmt.Sprintf("found no proposal_modules for %v", cc.ContractAddress))
		}
		log.Sugar.Errorf("error querying /list_proposals (%v): %v", resp.StatusCode, resp.Body)
		return nil, errors.New("error querying /list_proposals")
	}

	var proposals types.ProposalList
	body, err := ioutil.ReadAll(resp.Body)
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
			return nil, errors.New(fmt.Sprintf("could not decode proposals of contract %v: %v", cc.ContractAddress, resp.Body))
		}
		return propsList, nil
	}

	return &proposals, nil
}

func (cc *ContractClient) proposalModules(jsonData []byte) []string {
	resp, err := http.Post("http://localhost:8080/proposal_modules", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Sugar.Fatal(err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		log.Sugar.Errorf("error querying /proposal_modules (%v): %v", resp.StatusCode, resp.Body)
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

	resp, err := http.Post("http://localhost:8080/proposal", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Sugar.Fatal(err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Sugar.Errorf("error querying /proposal (%v): %v", resp.StatusCode, resp.Body)
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

func contracts() []string {
	file, err := os.Open("../contracts.txt")
	if err != nil {
		log.Sugar.Error(err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	sc := bufio.NewScanner(file)
	contracts := make([]string, 0)

	// Read through 'contracts' until an EOF is encountered.
	for sc.Scan() {
		contracts = append(contracts, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Sugar.Error(err)
	}
	return contracts
}

func UpdateContracts(cm *database.ContractManager, pm *database.ProposalManager) {
	for _, contractAddr := range contracts() {
		client := NewContractClient(contractAddr)
		config, err := client.config()
		if err != nil {
			log.Sugar.Errorf("while getting config for contract %v: %v", contractAddr, err)
			continue
		}
		proposals, err := client.proposals()
		if err != nil {
			log.Sugar.Errorf("while getting proposals for contract %v: %v", contractAddr, err)
			continue
		}

		contract, _ := cm.CreateOrUpdate(contractAddr, config)
		for _, proposal := range proposals.Proposals {
			dbProp, status := pm.CreateOrUpdate(contract, &proposal)
			if status == database.ProposalStatusChanged {
				log.Sugar.Infof("Proposal %v changed status to %v", dbProp.ID, dbProp.Status)
				//TODO: send notifications to users
			}
		}
		log.Sugar.Infof("updated contract %v (%v)", config.Name, contractAddr)
	}
}
