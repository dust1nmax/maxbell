package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

// 本项目使用简化版的投票分数
// 投一票就加432分   86400/200  --> 200张赞成票可以给你的帖子续一天

/* 投票的几种情况：
direction=1时，有两种情况：
	1. 之前没有投过票，现在投赞成票    --> 更新分数和投票记录
	2. 之前投反对票，现在改投赞成票    --> 更新分数和投票记录
direction=0时，有两种情况：
	1. 之前投过赞成票，现在要取消投票  --> 更新分数和投票记录
	2. 之前投过反对票，现在要取消投票  --> 更新分数和投票记录
direction=-1时，有两种情况：
	1. 之前没有投过票，现在投反对票    --> 更新分数和投票记录
	2. 之前投赞成票，现在改投反对票    --> 更新分数和投票记录

投票的限制：
每个贴子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票了。
	1. 到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
	2. 到期之后删除那个 KeyPostVotedZSetPF
*/

const (
	aweektime = 7 * 24 * 3600
	oncescore = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

func CreatePost(postid int64) error {
	//创建事务
	pipeline := rdb.TxPipeline()

	pipeline.ZAdd(GetKey(KeyPostTime), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postid,
	})

	pipeline.ZAdd(GetKey(KeyPostScore), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postid,
	})
	_, err := pipeline.Exec()
	return err
}

func PostVote(userid, postid string, value float64) error {
	var dir float64

	// 判断是否过期
	posttime := rdb.ZScore(GetKey(KeyPostTime), postid).Val()
	if float64(time.Now().Unix())-posttime > aweektime {
		return ErrVoteTimeExpire
	}
	// 更新帖子分数
	//创建事务
	pipeline := rdb.TxPipeline()
	rawvalue := rdb.ZScore(GetKey(KeyPostVotedPrefix+postid), userid).Val()
	if value > rawvalue {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(rawvalue - value)
	pipeline.ZIncrBy(GetKey(KeyPostScore), dir*diff*oncescore, postid)
	// 记录用户为帖子投票
	if value == 0 {
		pipeline.ZRem(GetKey(KeyPostVotedPrefix+userid), postid)
	} else {
		pipeline.ZAdd(GetKey(KeyPostVotedPrefix+userid), redis.Z{
			Score:  value,
			Member: userid,
		})
	}
	_, err := pipeline.Exec()
	return err
}
