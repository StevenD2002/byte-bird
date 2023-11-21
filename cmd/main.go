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

	_ "github.com/lib/pq"
	// "byte-bird/pkg/errors"
)

const (
	dbConnectionString    = "host=localhost port=5432 dbname=mydatabase user=postgres password=password sslmode=disable"
	redisConnectionString = "localhost:6379"
)

func main() {
	err := db.InitDB(dbConnectionString)
	if err != nil {
		log.Fatal(err)
	}


	userRepository := repository.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepository)

	redisPostRepository := repository.NewRedisPostRepository(redisConnectionString)
	postService := service.NewPostServiceImpl(redisPostRepository)

	httpServer := httpserver.NewHTTPServer(userService, postService)
	httpServer.StartServer()
}


