package logic

import (
	"maxbell/dao/redis"
	"maxbell/models"
	"strconv"

	"go.uber.org/zap"
)



func PostVote(userid int64, p *models.ParamPostVoteData) error {
	zap.L().Debug(
		"PostVote", 
		zap.Int64("userID", userid), 
		zap.String("postid", p.PostId), 
		zap.Int8("direction", p.Direction))
	return redis.PostVote(strconv.Itoa(int(userid)), p.PostId, float64(p.Direction))
}
