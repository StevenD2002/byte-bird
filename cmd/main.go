package main

import (
	// "database/sql"
	// "fmt"
	"byte-bird/internal/db"
	"byte-bird/internal/repository"
	"byte-bird/internal/service"
	"byte-bird/pkg/httpserver"
	"log"

	// "net/http"

	// _ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	// "byte-bird/pkg/errors"
)

const (
	// will change this later and throw stuff in an env
	dbConnectionString = "file:data.db?cache=shared&mode=rwc"
)

func main() {
	// Use the SQLite initialization function
	sqliteDB, err := db.InitSQLiteDB()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository(sqliteDB)
	userService := service.NewUserService(userRepository)

	postRepository := repository.NewPostRepository(sqliteDB)
	postService := service.NewPostService(postRepository)

	httpServer := httpserver.NewHTTPServer(userService, postService)
	httpServer.StartServer()
}
