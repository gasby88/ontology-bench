package main

import (
	"fmt"
	"github.com/ontio/ontology/account"
	"math/big"
	"time"
)

func TestTransfer() {
	taskCh := make(chan int, 1)
	doneCh := make(chan interface{}, 0)
	work := func() {
		for {
			select {
			case <-doneCh:
				return
			case t := <-taskCh:
				if t == 0 {
					close(doneCh)
					return
				}
				toAcc := account.NewAccount("")
				_, err := OntSdk.Rpc.Transfer("ont", Admin, toAcc, new(big.Int).SetInt64(1))
				if err != nil {
					fmt.Printf("Transfer error:%s\n", err)
					return
				}
			}
		}
	}

	for i := 0; i < WORKER_NUM; i++ {
		go work()
	}

	reqCount := 0
	timer := time.NewTicker(time.Second)
	for {
		select {
		case <-doneCh:
			return
		case <-timer.C:
			for i := 0; i < REQ_PER_SEC; i++ {
				taskCh <- 1
				reqCount++
				if reqCount == REQ_NUM {
					taskCh <- 0
				}
			}
		}
	}
}
