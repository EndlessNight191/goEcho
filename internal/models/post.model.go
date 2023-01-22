package models

import "database/sql"

func getPostById(db *sql.DB, request string) sql.Result {
	result, err := db.Exec("SELECT" + string(request))
	if err != nil {

	}
	return result
}

func getPosts(db *sql.DB, request string) sql.Result {
	result, err := db.Exec("SELECT" + string(request))
	if err != nil {

	}
	return result
}

func createPost(db *sql.DB, request string) sql.Result {
	result, err := db.Exec("INSERT" + string(request))
	if err != nil {

	}
	return result
}

func updatePost(db *sql.DB, request string) sql.Result {
	result, err := db.Exec("UPDATE" + string(request))
	if err != nil {

	}
	return result
}

func deletePost(db *sql.DB, request string) sql.Result {
	result, err := db.Exec("DELETE" + string(request))
	if err != nil {

	}
	return result
}
