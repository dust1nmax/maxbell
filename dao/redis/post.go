package redis

import (

	"maxbell/models"

	"github.com/go-redis/redis"
)

func GetPostIdsInOrder(p *models.ParamPostList) ([]string, error) {
	key := GetKey(KeyPostTime)
	if p.Order == models.OrderScore {
		key = GetKey(KeyPostScore)
	}
	start := (p.Page - 1) * p.Size
	stop := start + p.Size - 1
	return rdb.ZRevRange(key, start, stop).Result()
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
	for _, id := range ids{
		key := GetKey(KeyPostVotedPrefix) + id
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders{
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}
