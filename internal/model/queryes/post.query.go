package queryes

import (
	_ "database/sql"
	"goEcho/internal/database"
	"goEcho/internal/model/structs"

	"github.com/pingcap/errors"
)

func GetPostById(id int) (*structs.Post, error) {
	var post structs.Post

	if err := database.DB.QueryRow(`SELECT
	id, title, content, author_id, image 
	FROM posts(id) VALUES($1)`,
		id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.AuthorId,
		&post.Image,
	); err != nil {
		return nil, errors.Wrap(err, "QueryRow")
	}

	return &post, nil
}

func FindAllPost(limit int, offset int) ([]structs.Post, error) {
	var posts []structs.Post
	posts = make([]structs.Post, 0, limit)

	rows, err := database.DB.Query("SELECT id, title, content, author_id, image  FROM posts OFFSET $1 LIMIT $2", limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post structs.Post

		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Image,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func CreatePost(post *structs.Post) (*structs.Post, error) {
	if _, err := database.DB.Exec(
		"insert into posts (title, content, authorId, image) values ($1, $2, $3, $4)",
		post.Title, post.Content, post.AuthorId, post.Image); err != nil {
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
