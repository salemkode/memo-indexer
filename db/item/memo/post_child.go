package memo

import (
	"context"
	"github.com/jchavannes/jgo/jerr"
	"github.com/jchavannes/jgo/jutil"
	"github.com/memocash/index/db/client"
	"github.com/memocash/index/db/item/db"
	"github.com/memocash/index/ref/bitcoin/memo"
	"github.com/memocash/index/ref/config"
)

type PostChild struct {
	PostTxHash  []byte
	ChildTxHash []byte
}

func (c PostChild) GetUid() []byte {
	return jutil.CombineBytes(
		jutil.ByteReverse(c.PostTxHash),
		jutil.ByteReverse(c.ChildTxHash),
	)
}

func (c PostChild) GetShard() uint {
	return client.GetByteShard(c.PostTxHash)
}

func (c PostChild) GetTopic() string {
	return db.TopicMemoPostChild
}

func (c PostChild) Serialize() []byte {
	return nil
}

func (c *PostChild) SetUid(uid []byte) {
	if len(uid) != memo.TxHashLength*2 {
		return
	}
	c.PostTxHash = jutil.ByteReverse(uid[:32])
	c.ChildTxHash = jutil.ByteReverse(uid[32:])
}

func (c *PostChild) Deserialize([]byte) {}

func GetPostChildren(ctx context.Context, postTxHash []byte) ([]*PostChild, error) {
	shardConfig := config.GetShardConfig(db.GetShardByte32(postTxHash), config.GetQueueShards())
	dbClient := client.NewClient(shardConfig.GetHost())
	if err := dbClient.GetWOpts(client.Opts{
		Context:  ctx,
		Topic:    db.TopicMemoPostChild,
		Prefixes: [][]byte{jutil.ByteReverse(postTxHash)},
	}); err != nil {
		return nil, jerr.Get("error getting client message memo post children", err)
	}
	var postChildren = make([]*PostChild, len(dbClient.Messages))
	for i := range dbClient.Messages {
		postChildren[i] = new(PostChild)
		db.Set(postChildren[i], dbClient.Messages[i])
	}
	return postChildren, nil
}