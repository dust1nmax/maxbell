package logic

import (
	"maxbell/dao/mysql"
	"maxbell/models"
	"maxbell/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	//生成post id
	p.Id = snowflake.GenID()
	// 存到数据库
	return mysql.CreatePost(p)
	// 返回
}

// GetPostDetial 根据帖子id查询数据
func GetPostDetial(id int64) (data *models.ApiPostDetail, err error) {
	
	post, err := mysql.GetPostDetial(id)
	if err != nil {
		zap.L().Error("mysql.GetPostDetial(id) failed",
			zap.Int64("id", id),
			zap.Error(err))
		return
	}

	user, err := mysql.GetNameById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetNameById(id) failed",
			zap.Int64("user_id", post.AuthorID),
			zap.Error(err))
		return
	}

	community, err := mysql.GetCommunityDetialbyid(id)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetialbyid(id) failed",
			zap.Int64("community_id", post.CommunityID),
			zap.Error(err))
		return
	}

	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return

}
