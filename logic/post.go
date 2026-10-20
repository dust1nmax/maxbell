package logic

import (
	"maxbell/dao/mysql"
	"maxbell/models"
	"maxbell/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	//生成post id
	p.Id = snowflake.GenID()
	// 存到数据库
	return mysql.CreatePost(p)
	// 返回
}

// GetPostDetial 根据帖子id查询数据
func GetPostDetial(id int64) (data *models.Post, err error) {
	return mysql.GetPostDetial(id)
}
