package main

import (
	"log"
	"math/big"
	"os"
	"private/opsinTest/goAdmin/genFiles"
	"private/opsinTest/goAdmin/settings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/yaml.v2"
)

func main() {
	f, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(f, &settings.Cfg); err != nil {
		log.Fatal(err)
	}

	log.Println("get proposal info...")

	client, err := ethclient.Dial(settings.Cfg.SepoliaRPC)
	if err != nil {
		log.Fatal(err)
	}

	instance, err := genFiles.NewMyGovernor(common.HexToAddress(settings.Cfg.GovernorAddress), client)
	if err != nil {
		log.Fatal(err)
	}

	proposalID := new(big.Int)
	proposalID.SetString(settings.Cfg.ProposalID, 10)

	pp, err := instance.ProposalProposer(nil, proposalID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("proposal proposer:", pp)

	pd, err := instance.ProposalDeadline(nil, proposalID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("proposal deadline:", pd)

	pe, err := instance.ProposalEta(nil, proposalID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("proposal eta:", pe)

	vd, err := instance.ProposalVotes(nil, proposalID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("for votes:", vd.ForVotes, " against votes:", vd.AgainstVotes, " abstain votes:", vd.AbstainVotes)
}
