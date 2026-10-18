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
