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
	// 根据id查询post
	post, err := mysql.GetPostDetial(id)
	if err != nil {
		zap.L().Error("mysql.GetPostDetial(id) failed",
			zap.Int64("id", id),
			zap.Error(err))
		return
	}

	// 根据AuthorID查询username
	user, err := mysql.GetNameById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetNameById(post.AuthorID) failed",
			zap.Int64("user_id", post.AuthorID),
			zap.Error(err))
		return
	}

	// 根据CommunityID查询community
	community, err := mysql.GetCommunityDetialbyid(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetialbyid(post.CommunityID) failed",
			zap.Int64("community_id", post.CommunityID),
			zap.Error(err))
		return
	}

	// 拼接完整信息
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return

}

func GetPostList(page, size int64) (datas []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}

	datas = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts{
		// 根据AuthorID查询username
		user, err := mysql.GetNameById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetNameById(post.AuthorID) failed",
				zap.Int64("user_id", post.AuthorID),
				zap.Error(err))
			continue
		}

		// 根据CommunityID查询community
		community, err := mysql.GetCommunityDetialbyid(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetialbyid(post.CommunityID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}

		// 拼接完整信息
		data := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}

		datas = append(datas, data)
	}

	return
}
