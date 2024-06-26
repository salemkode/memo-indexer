package memo

import (
	"context"
	"fmt"
	"github.com/jchavannes/jgo/jutil"
	"github.com/memocash/index/db/client"
	"github.com/memocash/index/db/item/db"
	"github.com/memocash/index/ref/bitcoin/memo"
	"github.com/memocash/index/ref/config"
	"time"
)

type AddrName struct {
	Addr   [25]byte
	Seen   time.Time
	TxHash [32]byte
	Name   string
}

func (n *AddrName) GetUid() []byte {
	return jutil.CombineBytes(
		n.Addr[:],
		jutil.GetTimeByteNanoBig(n.Seen),
		jutil.ByteReverse(n.TxHash[:]),
	)
}

func (n *AddrName) GetShardSource() uint {
	return client.GenShardSource(n.Addr[:])
}

func (n *AddrName) GetTopic() string {
	return db.TopicMemoAddrName
}

func (n *AddrName) Serialize() []byte {
	return []byte(n.Name)
}

func (n *AddrName) SetUid(uid []byte) {
	if len(uid) != memo.AddressLength+memo.Int8Size+memo.TxHashLength {
		return
	}
	copy(n.Addr[:], uid[:25])
	n.Seen = jutil.GetByteTimeNanoBig(uid[25:33])
	copy(n.TxHash[:], jutil.ByteReverse(uid[33:65]))
}

func (n *AddrName) Deserialize(data []byte) {
	n.Name = string(data)
}

func GetAddrNames(ctx context.Context, addrs [][25]byte) ([]*AddrName, error) {
	var shardPrefixes = make(map[uint32][][]byte)
	for i := range addrs {
		shard := db.GetShardIdFromByte32(addrs[i][:])
		shardPrefixes[shard] = append(shardPrefixes[shard], addrs[i][:])
	}
	shardConfigs := config.GetQueueShards()
	var addrNames []*AddrName
	for shard, prefixes := range shardPrefixes {
		shardConfig := config.GetShardConfig(shard, shardConfigs)
		dbClient := client.NewClient(shardConfig.GetHost())
		if err := dbClient.GetWOpts(client.Opts{
			Topic:    db.TopicMemoAddrName,
			Prefixes: prefixes,
			Max:      1,
			Newest:   true,
			Context:  ctx,
		}); err != nil {
			return nil, fmt.Errorf("error getting db addr memo names by prefix; %w", err)
		}
		for _, msg := range dbClient.Messages {
			var addrName = new(AddrName)
			db.Set(addrName, msg)
			addrNames = append(addrNames, addrName)
		}
	}
	return addrNames, nil
}

func ListenAddrNames(ctx context.Context, addrs [][25]byte) (chan *AddrName, error) {
	var shardPrefixes = make(map[uint32][][]byte)
	for i := range addrs {
		shard := db.GetShardIdFromByte32(addrs[i][:])
		shardPrefixes[shard] = append(shardPrefixes[shard], addrs[i][:])
	}
	chanMessages, err := db.ListenPrefixes(ctx, db.TopicMemoAddrName, shardPrefixes)
	if err != nil {
		return nil, fmt.Errorf("error getting listen prefixes for memo addr names; %w", err)
	}
	var addrNameChan = make(chan *AddrName)
	go func() {
		for {
			msg, ok := <-chanMessages
			if !ok {
				return
			}
			var addrProfile = new(AddrName)
			db.Set(addrProfile, *msg)
			addrNameChan <- addrProfile
		}
	}()
	return addrNameChan, nil
}
