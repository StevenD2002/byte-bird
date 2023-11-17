package main

import (
	// "database/sql"
	// "fmt"
	"log"
	// "net/http"

	_ "github.com/lib/pq"
	"byte-bird/internal/db"
	"byte-bird/internal/repository"
	"byte-bird/internal/service"

	// "byte-bird/pkg/errors"
	"byte-bird/pkg/httpserver"
)

const (
	dbConnectionString = "host=localhost port=5432 dbname=mydatabase user=postgres password=password sslmode=disable"
)

func main() {
	err := db.InitDB(dbConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepository)

	httpServer := httpserver.NewHTTPServer(userService)
	httpServer.StartServer()
}
