package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"os"
	"private/opsinTest/goAdmin/genFiles"
	"private/opsinTest/goAdmin/settings"

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

	log.Println("deploying GovToken...")

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

	GovTokenAddress, tx, _, err := genFiles.DeployGovToken(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("GovTokenAddress:", GovTokenAddress)
	log.Println("deployment tx:", tx.Hash())

	log.Println("deploying TimeLock...")
	auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
	TimeLockAddress, tx, _, err := genFiles.DeployTimeLock(auth, client, big.NewInt(1), []common.Address{}, []common.Address{}, common.HexToAddress(settings.Cfg.AdminAddress))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("TimeLockAddress:", TimeLockAddress)
	log.Println("deployment tx:", tx.Hash())

	log.Println("deploying MyGovernor...")
	auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
	MyGovernorAddress, tx, _, err := genFiles.DeployMyGovernor(auth, client, GovTokenAddress, TimeLockAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("MyGovernorAddress:", MyGovernorAddress)
	log.Println("deployment tx:", tx.Hash())
}
