package redis

import (
	"maxbell/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func GetidsformKey(key string, page, size int64)([]string, error){
	start := (page - 1) * size
	stop := start + size- 1
	return rdb.ZRevRange(key, start, stop).Result()
}

func GetPostIdsInOrder(p *models.ParamPostList) ([]string, error) {
	key := GetKey(KeyPostTime)
	if p.Order == models.OrderScore {
		key = GetKey(KeyPostScore)
	}
	return GetidsformKey(key, p.Page, p.Size)
}

// GetPostVote 根据ids 查询每篇帖子的票数
func GetPostVote(ids []string) (data []int64, err error) {
	// data = make([]int64, 0, len(ids))
	// for _, id := range ids {
	// 	key := GetKey(KeyPostVotedPrefix) + id
	// 	v1 := rdb.ZCount(key, "1", "1").Val()
	// 	data = append(data, v1)
	// }

	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := GetKey(KeyPostVotedPrefix) + id
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

func GetCommunityPostIdsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	orderkey := GetKey(KeyPostTime)
	if p.Order == models.OrderScore {
		orderkey = GetKey(KeyPostScore)
	}
	// 使用zinterstore 把分区的帖子set与帖子分数的zset取交集生成一个新的zset
	// 使用缓存key减少使用zinterstore执行的次数
	key := orderkey + strconv.Itoa(int(p.CommunityId))
	comm_key := GetKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityId)))
	if rdb.Exists(orderkey).Val() < 1 {
		//不存在 就需要计算
		pipeline := rdb.Pipeline()

		pipeline.ZInterStore(
			key,
			redis.ZStore{
				Aggregate: "MAX",
			},
			comm_key,
			orderkey,
		)
		pipeline.Expire(key, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	return GetidsformKey(key, p.Page, p.Size)
}
