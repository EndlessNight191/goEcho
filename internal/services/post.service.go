package services

import (
	"goEcho/internal/model/queryes"
	"goEcho/internal/model/structs"
)

func GetPostById(id int) (*structs.Post, error) {
	post, err := queryes.GetPostById(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func GetPosts(limit int, offset int) ([]structs.Post, error) {
	posts, err := queryes.FindAllPost(limit, offset)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func CreatePost(post *structs.Post) (*structs.Post, error) {
	post, err := queryes.CreatePost(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func UpdatePostById(id int, postUpdate *structs.Post) (*structs.Post, error) {
	post, err := GetPostById(id)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, nil
	}

	resultUpdate, err := queryes.UpdatePostById(id, postUpdate)
	if err != nil {
		return nil, err
	}
	return resultUpdate, nil
}

func DeletePostById(id int) error {
	post, err := GetPostById(id)
	if err != nil || post == nil {
		return err
	}

	err = queryes.DeleteOnePostById(id)
	if err != nil {
		return err
	}
	return nil
}
