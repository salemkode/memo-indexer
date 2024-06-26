package saver

import (
	"context"
	"fmt"
	"github.com/jchavannes/btcd/chaincfg/chainhash"
	"github.com/memocash/index/db/item/chain"
	"github.com/memocash/index/db/item/db"
	"github.com/memocash/index/ref/dbi"
	"log"
)

type TxMinimal struct {
	Verbose bool
}

func (t *TxMinimal) SaveTxs(ctx context.Context, block *dbi.Block) error {
	if block.IsNil() {
		return fmt.Errorf("error nil block")
	}
	if err := t.QueueTxs(block); err != nil {
		return fmt.Errorf("error queueing tx minimal block; %w", err)
	}
	return nil
}

func (t *TxMinimal) QueueTxs(block *dbi.Block) error {
	blockHash := block.Header.BlockHash()
	var objects []db.Object
	for _, dbiTx := range block.Transactions {
		tx := dbiTx.MsgTx
		txHash := chainhash.Hash(dbiTx.Hash)
		if t.Verbose {
			log.Printf("tx: %s\n", txHash.String())
		}
		if block.HasHeader() {
			var blockTx = &chain.BlockTx{
				BlockHash: blockHash,
				Index:     dbiTx.BlockIndex,
				TxHash:    txHash,
			}
			objects = append(objects, blockTx)
			objects = append(objects, &chain.TxBlock{
				TxHash:    txHash,
				BlockHash: blockHash,
				Index:     dbiTx.BlockIndex,
			})
		}
		if dbiTx.Saved {
			continue
		}
		objects = append(objects, &chain.Tx{
			TxHash:   txHash,
			Version:  tx.Version,
			LockTime: tx.LockTime,
		})
		for j := range tx.TxIn {
			objects = append(objects, &chain.TxInput{
				TxHash:       txHash,
				Index:        uint32(j),
				PrevHash:     tx.TxIn[j].PreviousOutPoint.Hash,
				PrevIndex:    tx.TxIn[j].PreviousOutPoint.Index,
				Sequence:     tx.TxIn[j].Sequence,
				UnlockScript: tx.TxIn[j].SignatureScript,
			})
			objects = append(objects, &chain.OutputInput{
				PrevHash:  tx.TxIn[j].PreviousOutPoint.Hash,
				PrevIndex: tx.TxIn[j].PreviousOutPoint.Index,
				Hash:      txHash,
				Index:     uint32(j),
			})
		}
		for k := range tx.TxOut {
			objects = append(objects, &chain.TxOutput{
				TxHash:     txHash,
				Index:      uint32(k),
				Value:      tx.TxOut[k].Value,
				LockScript: tx.TxOut[k].PkScript,
			})
		}
		objects = append(objects, &chain.TxSeen{
			TxHash:    txHash,
			Timestamp: dbiTx.Seen,
		})
	}
	if err := db.Save(objects); err != nil {
		return fmt.Errorf("error saving db tx objects; %w", err)
	}
	return nil
}

func NewTxMinimal(verbose bool) *TxMinimal {
	return &TxMinimal{
		Verbose: verbose,
	}
}
