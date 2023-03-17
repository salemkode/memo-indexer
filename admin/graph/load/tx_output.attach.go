package load

import (
	"fmt"
	"github.com/memocash/index/admin/graph/model"
	"github.com/memocash/index/db/item/chain"
	"github.com/memocash/index/ref/bitcoin/memo"
)

type Outputs struct {
	baseA
	Outputs []*model.TxOutput
}

func AttachToOutputs(preloads []string, outputs []*model.TxOutput) error {
	if len(outputs) == 0 {
		return nil
	}
	o := Outputs{
		baseA:   baseA{Preloads: preloads},
		Outputs: outputs,
	}
	o.Wait.Add(2)
	go o.AttachInfo()
	go o.AttachSpends()
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
	defer o.Wait.Done()
	if !o.HasPreload([]string{"amount", "script"}) {
		return
	}
	outs := o.GetOuts(true)
	if len(outs) == 0 {
		return
	}
	txOutputs, err := chain.GetTxOutputs(outs)
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
	if !o.HasPreload([]string{"spends"}) {
		return
	}
	outs := o.GetOuts(false)
	spends, err := chain.GetOutputInputs(outs)
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
	preloads := GetPrefixPreloads(o.Preloads, "spends.")
	o.Mutex.Unlock()
	if err := AttachToInputs(preloads, allSpends); err != nil {
		o.AddError(fmt.Errorf("error attaching to tx inputs spends for model tx outputs; %w", err))
		return
	}
}