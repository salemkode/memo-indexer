package attach

import (
	"context"
	"fmt"
	"github.com/memocash/index/db/item/chain"
	"github.com/memocash/index/db/item/slp"
	"github.com/memocash/index/graph/model"
	"github.com/memocash/index/ref/bitcoin/memo"
	"github.com/memocash/index/ref/bitcoin/wallet"
	"sync"
)

type Outputs struct {
	base
	Outputs     []*model.TxOutput
	DetailsWait sync.WaitGroup
}

func ToOutputs(ctx context.Context, fields Fields, outputs []*model.TxOutput) error {
	if len(outputs) == 0 {
		return nil
	}
	o := Outputs{
		base:    base{Ctx: ctx, Fields: fields},
		Outputs: outputs,
	}
	o.DetailsWait.Add(1)
	o.Wait.Add(5)
	go o.AttachInfo()
	go o.AttachSpends()
	go o.AttachSlps()
	go o.AttachSlpBatons()
	go o.AttachTxs()
	o.DetailsWait.Wait()
	go o.AttachLocks()
	o.Wait.Wait()
	if len(o.Errors) > 0 {
		return fmt.Errorf("error attaching to outputs; %w", o.Errors[0])
	}
	return nil
}

func (o *Outputs) GetOuts(checkScript bool) []memo.Out {
	o.Mutex.Lock()
	defer o.Mutex.Unlock()
	var outs []memo.Out
	for i := range o.Outputs {
		if checkScript && len(o.Outputs[i].Script) != 0 {
			continue
		}
		outs = append(outs, memo.Out{
			TxHash: o.Outputs[i].Hash[:],
			Index:  o.Outputs[i].Index,
		})
	}
	return outs
}

func (o *Outputs) AttachInfo() {
	defer o.DetailsWait.Done()
	if !o.HasField([]string{"amount", "script", "lock"}) {
		return
	}
	outs := o.GetOuts(true)
	if len(outs) == 0 {
		return
	}
	txOutputs, err := chain.GetTxOutputs(o.Ctx, outs)
	if err != nil {
		o.AddError(fmt.Errorf("error getting tx outputs for model tx; %w", err))
		return
	}
	o.Mutex.Lock()
	defer o.Mutex.Unlock()
	for i := range o.Outputs {
		for j := range txOutputs {
			if o.Outputs[i].Hash != txOutputs[j].TxHash || o.Outputs[i].Index != txOutputs[j].Index {
				continue
			}
			o.Outputs[i].Amount = txOutputs[j].Value
			o.Outputs[i].Script = txOutputs[j].LockScript
			break
		}
	}
}

func (o *Outputs) AttachSpends() {
	defer o.Wait.Done()
	if !o.HasField([]string{"spends"}) {
		return
	}
	outs := o.GetOuts(false)
	spends, err := chain.GetOutputInputs(o.Ctx, outs)
	if err != nil {
		o.AddError(fmt.Errorf("error getting tx inputs spends for model tx outputs; %w", err))
		return
	}
	var allSpends []*model.TxInput
	o.Mutex.Lock()
	for i := range o.Outputs {
		for j := range spends {
			if o.Outputs[i].Hash != spends[j].PrevHash || o.Outputs[i].Index != spends[j].PrevIndex {
				continue
			}
			o.Outputs[i].Spends = append(o.Outputs[i].Spends, &model.TxInput{
				Hash:      spends[j].Hash,
				Index:     spends[j].Index,
				PrevHash:  spends[j].PrevHash,
				PrevIndex: spends[j].PrevIndex,
			})
		}
		allSpends = append(allSpends, o.Outputs[i].Spends...)
	}
	o.Mutex.Unlock()
	if err := ToInputs(o.Ctx, GetPrefixFields(o.Fields, "spends."), allSpends); err != nil {
		o.AddError(fmt.Errorf("error attaching to tx inputs spends for model tx outputs; %w", err))
		return
	}
}

func (o *Outputs) AttachSlps() {
	defer o.Wait.Done()
	if !o.HasField([]string{"slp"}) {
		return
	}
	outs := o.GetOuts(false)
	slpOutputs, err := slp.GetOutputs(o.Ctx, outs)
	if err != nil {
		o.AddError(fmt.Errorf("error getting slp outputs for model tx outputs; %w", err))
		return
	}
	var allSlpOutputs []*model.SlpOutput
	o.Mutex.Lock()
	for i := range o.Outputs {
		for j := range slpOutputs {
			if o.Outputs[i].Hash != slpOutputs[j].TxHash || o.Outputs[i].Index != slpOutputs[j].Index {
				continue
			}
			o.Outputs[i].Slp = &model.SlpOutput{
				Hash:      slpOutputs[j].TxHash,
				Index:     slpOutputs[j].Index,
				TokenHash: slpOutputs[j].TokenHash,
				Amount:    slpOutputs[j].Quantity,
			}
			allSlpOutputs = append(allSlpOutputs, o.Outputs[i].Slp)
		}
	}
	o.Mutex.Unlock()
	if err := ToSlpOutputs(o.Ctx, GetPrefixFields(o.Fields, "slp."), allSlpOutputs); err != nil {
		o.AddError(fmt.Errorf("error attaching to slp outputs for tx outputs; %w", err))
		return
	}
}

func (o *Outputs) AttachSlpBatons() {
	defer o.Wait.Done()
	if !o.HasField([]string{"slp_baton"}) {
		return
	}
	outs := o.GetOuts(false)
	slpBatons, err := slp.GetBatons(o.Ctx, outs)
	if err != nil {
		o.AddError(fmt.Errorf("error getting slp batons for model tx outputs; %w", err))
		return
	}
	var allSlpBatons []*model.SlpBaton
	o.Mutex.Lock()
	for i := range o.Outputs {
		for j := range slpBatons {
			if o.Outputs[i].Hash != slpBatons[j].TxHash || o.Outputs[i].Index != slpBatons[j].Index {
				continue
			}
			o.Outputs[i].SlpBaton = &model.SlpBaton{
				Hash:      slpBatons[j].TxHash,
				Index:     slpBatons[j].Index,
				TokenHash: slpBatons[j].TokenHash,
			}
			allSlpBatons = append(allSlpBatons, o.Outputs[i].SlpBaton)
		}
	}
	o.Mutex.Unlock()
	if err := ToSlpBatons(o.Ctx, GetPrefixFields(o.Fields, "slp_baton."), allSlpBatons); err != nil {
		o.AddError(fmt.Errorf("error attaching to slp batons for tx outputs; %w", err))
		return
	}
}

func (o *Outputs) AttachTxs() {
	defer o.Wait.Done()
	if !o.HasField([]string{"tx"}) {
		return
	}
	var allTxs []*model.Tx
	o.Mutex.Lock()
	for j := range o.Outputs {
		o.Outputs[j].Tx = &model.Tx{Hash: o.Outputs[j].Hash}
		allTxs = append(allTxs, o.Outputs[j].Tx)
	}
	o.Mutex.Unlock()
	if err := ToTxs(o.Ctx, GetPrefixFields(o.Fields, "tx."), allTxs); err != nil {
		o.AddError(fmt.Errorf("error attaching to txs for model tx outputs; %w", err))
		return
	}
}

func (o *Outputs) AttachLocks() {
	defer o.Wait.Done()
	if !o.HasField([]string{"lock"}) {
		return
	}
	var allLocks []*model.Lock
	o.Mutex.Lock()
	for j := range o.Outputs {
		address, err := wallet.GetAddrFromLockScript(o.Outputs[j].Script)
		if err != nil {
			continue
		}
		o.Outputs[j].Lock = &model.Lock{Address: model.Address(*address)}
		allLocks = append(allLocks, o.Outputs[j].Lock)
	}
	o.Mutex.Unlock()
	if len(allLocks) == 0 {
		return
	}
	if err := ToLocks(o.Ctx, GetPrefixFields(o.Fields, "lock."), allLocks); err != nil {
		o.AddError(fmt.Errorf("error attaching to locks for model tx outputs; %w", err))
		return
	}
}
