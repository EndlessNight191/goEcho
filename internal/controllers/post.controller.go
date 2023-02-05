package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"goEcho/internal/model/structs"
	"goEcho/internal/services"
)

func GetPostById(c echo.Context) error {
	id := c.Param("id")
	post, err := services.GetPostById(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, post)
}

func GetPosts(c echo.Context) error {
	return nil
}

func CreatePost(c echo.Context) error {
	u := new(structs.Post)
	if err := c.Bind(u); err != nil {
		return err
	}

	resultCreate, err := services.CreatePost(u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resultCreate)
}

func UpdatePostById(c echo.Context) error {
	return nil
}

func DeletePostById(c echo.Context) error {
	return nil
}
