package attach

import (
	"context"
	"fmt"
	"github.com/memocash/index/graph/model"
)

type MemoFollow struct {
	base
	Follows []*model.Follow
}

func ToMemoFollows(ctx context.Context, fields []Field, follows []*model.Follow) error {
	if len(follows) == 0 {
		return nil
	}
	o := MemoFollow{
		base:    base{Ctx: ctx, Fields: fields},
		Follows: follows,
	}
	o.Wait.Add(3)
	go o.AttachLocks()
	go o.AttachFollowLocks()
	go o.AttachTxs()
	o.Wait.Wait()
	if len(o.Errors) > 0 {
		return fmt.Errorf("error attaching to memo follows; %w", o.Errors[0])
	}
	return nil
}

func (a *MemoFollow) AttachLocks() {
	defer a.Wait.Done()
	if !a.HasField([]string{"lock"}) {
		return
	}
	var allLocks []*model.Lock
	a.Mutex.Lock()
	for _, follow := range a.Follows {
		follow.Lock = &model.Lock{Address: follow.Address}
		allLocks = append(allLocks, follow.Lock)
	}
	a.Mutex.Unlock()
	if err := ToLocks(a.Ctx, GetPrefixFields(a.Fields, "lock."), allLocks); err != nil {
		a.AddError(fmt.Errorf("error attaching to locks for memo follows; %w", err))
		return
	}
}

func (a *MemoFollow) AttachFollowLocks() {
	defer a.Wait.Done()
	if !a.HasField([]string{"follow_lock"}) {
		return
	}
	var allLocks []*model.Lock
	a.Mutex.Lock()
	for _, follow := range a.Follows {
		follow.FollowLock = &model.Lock{Address: follow.FollowAddress}
		allLocks = append(allLocks, follow.FollowLock)
	}
	a.Mutex.Unlock()
	if err := ToLocks(a.Ctx, GetPrefixFields(a.Fields, "follow_lock."), allLocks); err != nil {
		a.AddError(fmt.Errorf("error attaching to follow locks for memo follows; %w", err))
		return
	}
}

func (a *MemoFollow) AttachTxs() {
	defer a.Wait.Done()
	if !a.HasField([]string{"tx"}) {
		return
	}
	var allTxs []*model.Tx
	a.Mutex.Lock()
	for _, follow := range a.Follows {
		follow.Tx = &model.Tx{Hash: follow.TxHash}
		allTxs = append(allTxs, follow.Tx)
	}
	a.Mutex.Unlock()
	if err := ToTxs(a.Ctx, GetPrefixFields(a.Fields, "tx."), allTxs); err != nil {
		a.AddError(fmt.Errorf("error attaching to txs for memo follows; %w", err))
		return
	}
}
