package controllers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"

	"goEcho/internal/model/structs"
	"goEcho/internal/services"
)

type response struct {
	Message string `json:"message"`
}

func GetPostById(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
		log.Printf("error: %v", err)
		return c.JSON(http.StatusBadRequest, response{Message: "invalid id"})
	}

	post, err := services.GetPostById(ID)
	if err != nil {
		return err
		log.Printf("error: %v", err)
		return c.JSON(http.StatusBadRequest, response{Message: "invalid id"})
	}

	return c.JSON(http.StatusOK, post)
}

func GetPosts(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	limitS, err := strconv.Atoi(limit)
	if err != nil {
		log.Printf("error: %v", err)
		return c.JSON(http.StatusBadRequest, response{Message: "invalid limit"})
	}
	offsetS, err := strconv.Atoi(offset)
	if err != nil {
		log.Printf("error: %v", err)
		return c.JSON(http.StatusBadRequest, response{Message: "invalid offset"})
	}

	posts, err := services.GetPosts(int(limitS), int(offsetS))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, posts)
}

func CreatePost(c echo.Context) error {
	var a structs.Post

	if err := c.Bind(a); err != nil {
		log.Printf("error: %v", err)
		return c.JSON(http.StatusBadRequest, response{Message: "invalid post"})
	}

	resultCreate, err := services.CreatePost(&a)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resultCreate)
}

func UpdatePostById(c echo.Context) error {
	var a structs.Post
	id := c.Param("id")
	s, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return err
	}

	if err := c.Bind(a); err != nil {
		log.Printf("error: %v", err)
		return c.JSON(http.StatusBadRequest, response{Message: "invalid post"})
	}

	postUpdate, err := services.UpdatePostById(int(s), &a)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, postUpdate)
}

func DeletePostById(c echo.Context) error {
	id := c.QueryParam("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error: %v", err)
		return c.JSON(http.StatusBadRequest, response{Message: "invalid id"})
	}

	err = services.DeletePostById(ID)
	if err != nil {
		log.Printf("error: %v", err)
		return c.JSON(http.StatusBadRequest, response{Message: "error deleted"})
	}

	return c.JSON(http.StatusOK, "success")
}
