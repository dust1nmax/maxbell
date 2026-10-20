package mysql

import "maxbell/models"

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
	from post where post_id = ？`
	err = db.Get(post, sqlstr, id)
	return
}
