package services

import (
	"goEcho/internal/model/queryes"
	"goEcho/internal/model/structs"
)

func GetPostById(id string) (*structs.Post, error) {
	post, err := queryes.GetPostById(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func GetPosts(limit int, offset int) error {
	return nil
}

func CreatePost(post *structs.Post) (*structs.Post, error) {
	post, err := queryes.CreatePost(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func UpdatePostById(id int, postUpdate *structs.Post) (*structs.Post, error) {
	resultUpdate, err := queryes.UpdatePostById(id, postUpdate)
	if err != nil {
		return nil, err
	}
	return resultUpdate, nil
}

func DeletePostById(id int) error {
	err := queryes.DeleteOnePostById(id)
	if err != nil {
		return err
	}
	return nil
}
