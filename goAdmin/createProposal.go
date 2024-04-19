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

	log.Println("creating new propose...")

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

	//test values for creating new proposal, you can edit them here
	targets := []common.Address{common.HexToAddress("0x70B06c56Ef6697BDd4D3B1d064e952c386BCF0C0")}
	values := []*big.Int{big.NewInt(0)}
	calldatas := [][]byte{[]byte("0xb921e16300000000000000000000000000000000000000000000000000000000000003e8")}
	description := "testPropose"
	tx, err := instance.Propose(auth, targets, values, calldatas, description)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tx.Hash())

}
