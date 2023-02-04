package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"goEcho/internal/model/queryes"
	"goEcho/internal/model/structs"
	"goEcho/internal/services"
)

func GetPostById(c echo.Context) error {
	post, err := services.GetPostById(c.QueryParam("id"))
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
	result, err := queryes.CreatePost(u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func UpdatePostById(c echo.Context) error {
	return nil
}

func DeletePostById(c echo.Context) error {
	return nil
}
