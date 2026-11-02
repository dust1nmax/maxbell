package redis

import (
	"maxbell/models"

)

func GetPostIdsInOrder(p *models.ParamPostList)([]string ,error ){
	key := GetKey(KeyPostTime)
	if p.Order == models.OrderScore {
		key = GetKey(KeyPostScore)
	}
	start := (p.Page - 1) * p.Size
	stop := start + p.Size - 1
	return rdb.ZRevRange(key, start, stop).Result()
} 