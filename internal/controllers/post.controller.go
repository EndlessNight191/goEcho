package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"goEcho/internal/model/structs"
	"goEcho/internal/services"
)

func GetPostById(c echo.Context) error {
	id := c.Param("id") // переменная ID 2 большими буквами
	// s, err := strconv.ParseInt(id, 10, 32) // кал е баный
	// if err != nil {
	// 	return err
	// }

	type response struct {
		Message string `json:"message"`
	}

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("error: %v", err)

		return c.JSON(http.StatusBadRequest, response{Message: "invalid id"}) // почему в контроллкрк не возрващается хуевый статус + можно сделать структурку в которую ты буш запихивать сообщение
	}

	post, err := services.GetPostById(ID)
	if err != nil {
		return err // та же хуйня везде
	}

	return c.JSON(http.StatusOK, post)
}

func GetPosts(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	limitS, err := strconv.ParseInt(limit, 10, 32) // atoi
	if err != nil {
		return err // log ошибки плюс нормальный return
	}
	offsetS, err := strconv.ParseInt(offset, 10, 32) // ATOI
	if err != nil {
		return err // log ошибки плюс нормальный return
	}

	posts, err := services.GetPosts(int(limitS), int(offsetS))
	if err != nil {
		return err // log ошибки плюс нормальный return
	}

	return c.JSON(http.StatusOK, posts)
}

func CreatePost(c echo.Context) error {
	// u := new(structs.Post)// нахуя через new?)
	var a structs.Post

	if err := c.Bind(a); err != nil {
		return err // log ошибки плюс нормальный return
	}

	resultCreate, err := services.CreatePost(a)
	if err != nil {
		return err // log ошибки плюс нормальный return
	}
	return c.JSON(http.StatusOK, resultCreate)
}

func UpdatePostById(c echo.Context) error {
	ID := c.Param("id")
	s, err := strconv.ParseInt(id, 10, 32) // ATOI
	if err != nil {
		return err // log ошибки плюс нормальный return
	}

	u := new(structs.Post)
	if err := c.Bind(u); err != nil {
		return err // log ошибки плюс нормальный return
	}

	postUpdate, err := services.UpdatePostById(int(s), u)
	if err != nil {
		return err // log ошибки плюс нормальный return
	}

	return c.JSON(http.StatusOK, postUpdate)
}

func DeletePostById(c echo.Context) error {
	id := c.Param("id")
	s, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return err
	}

	if err := services.DeletePostById(int(s)); err != nil {
		return err
	} // тварь if можно юзать если не нужна переменная после if

	return c.JSON(http.StatusOK, "success")

}
