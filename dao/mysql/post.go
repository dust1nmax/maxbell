package mysql

import (
	"maxbell/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

func CreatePost(p *models.Post) (err error) {
	sqlstr := `insert into post(
		post_id, title, content, author_id, community_id)
		values (?, ?, ?, ?, ?)
		`
	_, err = db.Exec(sqlstr, p.Id, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostDetial(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlstr := `select 
	post_id, title, content, author_id, community_id, create_time
	from post where post_id = ?`
	err = db.Get(post, sqlstr, id)
	return
}

func GetNameById(id int64) (user *models.User, err error) {
	user = new(models.User)
	sqlstr := `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlstr, id)
	return
}

// GetPostList 每页有size个帖子 要查询第page页
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlstr := `select 
	post_id, title, content, author_id, community_id, create_time
	from post
	ORDER BY create_time
	DESC 
	limit ?,?` // LIMIT 起始偏移量, 查询数量

	posts = make([]*models.Post, 0, 2)

	// 查询第page页 就要先跳过 (page-1)*size 篇帖子
	err = db.Select(&posts, sqlstr, (page-1)*size, size)
	return
}

func GetPostsByIds(ids []string) (postlist []*models.Post, err error) {
	sqlstr := `select 
	post_id, title, content, author_id, community_id, create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id, ?)
	`
	query, args, err := sqlx.In(sqlstr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	db.Select(&postlist, query, args...)
	return
}
