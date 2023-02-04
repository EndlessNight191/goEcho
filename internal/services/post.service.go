package services

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"goEcho/internal/model/queryes"
	"goEcho/internal/model/structs"
)

func GetPostById(id string) (*sql.Rows, error) {
	post, err := queryes.GetPostById(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func GetPosts(c echo.Context) error {
	return nil
}

func CreatePost(post *structs.Post) (*structs.Post, error) {
	post, err := queryes.CreatePost(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func UpdatePostById() error {
	return nil
}

func DeletePostById() error {
	return nil
}
