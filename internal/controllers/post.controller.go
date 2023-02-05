package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"goEcho/internal/model/structs"
	"goEcho/internal/services"
)

func GetPostById(c echo.Context) error {
	id := c.Param("id")
	s, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return err
	}

	post, err := services.GetPostById(int(s))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, post)
}

func GetPosts(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	limitS, err := strconv.ParseInt(limit, 10, 32)
	if err != nil {
		return err
	}
	offsetS, err := strconv.ParseInt(offset, 10, 32)
	if err != nil {
		return err
	}

	posts, err := services.GetPosts(int(limitS), int(offsetS))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, posts)
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
	id := c.Param("id")
	s, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return err
	}

	u := new(structs.Post)
	if err := c.Bind(u); err != nil {
		return err
	}

	postUpdate, err := services.UpdatePostById(int(s), u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, postUpdate)
}

func DeletePostById(c echo.Context) error {
	id := c.Param("id")
	s, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return err
	}

	err = services.DeletePostById(int(s))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "success")

}
