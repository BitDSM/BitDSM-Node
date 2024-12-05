package main

import (
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"github.com/BitDSM/BitDSM-Node/BtcDepositConfirmer"
	"github.com/BitDSM/BitDSM-Node/api"
	"github.com/BitDSM/BitDSM-Node/ethComms"
	"github.com/BitDSM/BitDSM-Node/operator"
	"github.com/BitDSM/BitDSM-Node/utils"
)

func initialize() {
	utils.InitConfigFile()
	utils.LoadBtcWallet(viper.GetString("wallet_name"))
	ethAccount := ethComms.LoadEthAccount()
	fmt.Println("Eth account: ", ethAccount.Address.Hex())
	operator.RegisterOperator()
}

func main() {
	initialize()
	var wg sync.WaitGroup
	wg.Add(1)
	go api.Server()
	wg.Add(1)
	go BtcDepositConfirmer.CheckDeposit()
	wg.Add(1)
	go ethComms.SubscribeToDepositRequests()
	wg.Add(1)
	go BtcDepositConfirmer.CheckWithdraw()
	ethComms.SubscribeToWithdrawRequests()
	wg.Wait()
}

//abigen --bin=BitdsmRegistry/BitdsmRegistry.bin  --abi=BitdsmRegistry/BitdsmRegistry.abi  --pkg=BitdsmRegistry --out=BitdsmRegistry/BitdsmRegistry.go
