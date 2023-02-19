package main

import (
	"follme/comment-service/pkg/adapter/database"
	"follme/comment-service/pkg/server"
)

func main() {
	database.ConnectDB()
	server.Route()
}
