package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"os"
	"private/opsinTest/deployer/genFiles"
	"private/opsinTest/deployer/settings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

	log.Println("voting...")

	client, err := ethclient.Dial(settings.Cfg.SepoliaRPC)
	if err != nil {
		log.Fatal(err)
	}

	_privateKey := settings.Cfg.AdminPrivateKey
	privateKey, err := crypto.HexToECDSA(_privateKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal(errors.New("error casting public key to ECDSA"))
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(settings.Cfg.SepoliaChainID)))
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(5000000) // in units
	auth.GasPrice = gasPrice
	auth.From = fromAddress

	instance, err := genFiles.NewMyGovernor(common.HexToAddress(settings.Cfg.GovernorAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	proposalID := new(big.Int)

	proposalID.SetString(settings.Cfg.ProposalID, 10)

	//change vote type here
	support := 1
	tx, err := instance.CastVote(auth, proposalID, uint8(support))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tx.Hash())

}
