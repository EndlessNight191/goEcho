package queryes

import (
	"database/sql"
	"goEcho/internal/database"
	"goEcho/internal/model/structs"
)

func GetPostById(id string) (*sql.Rows, error) {
	post, err := database.DB.Query(
		"select * from posts (id) values ($1)",
		id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func FindAllPost(limit int, offset int) (*sql.Rows, error) {
	posts, err := database.DB.Query(
		"select * from posts OFFSET $1 ROWS FETCH NEXT $2 ROWS ONLY;",
		limit, offset)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func CreatePost(post *structs.Post) (*structs.Post, error) {
	_, err := database.DB.Exec(
		"insert into posts (title, content, authorId, image) values ($1, $2, $3, $4)",
		post.Title, post.Content, post.AuthorId, post.Image)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func UpdatePostById(id int, postUpdate *structs.Post) (*structs.Post, error) {
	_, err := database.DB.Exec(
		"update posts where id = $1 set title = $2, content = $3, image = $4",
		id, postUpdate.Title, postUpdate.Content, postUpdate.Image)
	if err != nil {
		return nil, err
	}
	return postUpdate, nil
}

func DeleteOnePostById(id int) error {
	_, err := database.DB.Exec(
		"delete posts where id = $1",
		id)
	if err != nil {
		return err
	}
	return nil
}
