package main

import (
	"encoding/hex"
	"example/common"
	"github.com/memocash/index/ref/bitcoin/memo"
	"github.com/memocash/index/ref/bitcoin/tx/gen"
	"github.com/memocash/index/ref/bitcoin/tx/parse"
	"github.com/memocash/index/ref/bitcoin/tx/script"
	"log"
	"os"
	"strings"
)

func main() {
	var msg = strings.Join(os.Args[1:], " ")
	if msg == "" {
		log.Fatalf("No memo post message provided.")
	}
	wlt, err := common.NewWalletFromStdinWif()
	if err != nil {
		log.Fatalf("error creating new wallet; %v", err)
	}
	tx, err := gen.Tx(gen.TxRequest{
		Getter: wlt,
		Outputs: []*memo.Output{{Script: &script.Post{
			Message: msg,
		}}},
		Change:  wlt.Change,
		KeyRing: wlt.KeyRing,
	})
	if err != nil {
		log.Fatalf("error generating memo post tx; %v", err)
	}
	txInfo := parse.GetTxInfo(tx)
	txInfo.Print()
	if err := wlt.Client.Broadcast(hex.EncodeToString(txInfo.Raw)); err != nil {
		log.Fatalf("error broadcasting tx; %v", err)
	}
	log.Println("Tx broadcast!")
}
