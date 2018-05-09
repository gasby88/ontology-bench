package main

import (
	"flag"
	"fmt"
	"github.com/ontio/ontology-crypto/keypair"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/account"
	"github.com/ontio/ontologybp/common/log"
	"math/big"
)

var (
	REQ_NUM     int
	REQ_PER_SEC int
	WORKER_NUM  int
	RPC_ADDRESS string
	WALLET_FILE string
	WALLET_PWD  string
)

var (
	OntSdk *sdk.OntologySdk
	Admin  *account.Account
)

func init() {
	flag.IntVar(&REQ_NUM, "r", 1000, "Request num")
	flag.IntVar(&REQ_PER_SEC, "tps", 100, "Request per second")
	flag.IntVar(&WORKER_NUM, "w", 10, "Worker num")
	flag.StringVar(&RPC_ADDRESS, "addr", "http://localhost:20336", "Address of ontology rpc")
	flag.StringVar(&WALLET_FILE, "wallet", "./wallet.dat", "Wallet file path")
	flag.StringVar(&WALLET_PWD, "pwd", "wangbing", "Password of wallet")
	flag.Parse()
}

func main() {
	log.Init()
	OntSdk = sdk.NewOntologySdk()
	OntSdk.Rpc.SetAddress(RPC_ADDRESS)
	wallet, err := OntSdk.OpenWallet(WALLET_FILE, WALLET_PWD)
	if err != nil {
		fmt.Printf("OpenWallet error:%s\n", err)
		return
	}
	Admin, err = wallet.GetDefaultAccount()
	if err != nil {
		fmt.Printf("CreateAccount error:%s", err)
		return
	}
	fmt.Printf("Admin:%x\n", keypair.SerializePublicKey(Admin.PublicKey))

	balance, err := OntSdk.Rpc.GetBalance(Admin.Address)
	if err != nil {
		fmt.Printf("GetBalance error:%s\n", err)
		return
	}

	fmt.Printf("Admin ont balance:%d\n", balance.Ont)
	if balance.Ont.Cmp(new(big.Int)) == 0 {
		fmt.Printf("Admin balance not enought\n")
		return
	}
	TestTransfer()
}
