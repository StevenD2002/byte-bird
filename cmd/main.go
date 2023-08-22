package main

import (
    // "database/sql"
    // "fmt"
    "log"
    // "net/http"

    _ "github.com/lib/pq"
    "github.com/stevend2002/tgp-bp/internal/db"
    "github.com/stevend2002/tgp-bp/internal/repository"
    "github.com/stevend2002/tgp-bp/internal/service"
    // "github.com/stevend2002/tgp-bp/pkg/errors"
    "github.com/stevend2002/tgp-bp/pkg/httpserver"
)


const (
    dbConnectionString = "host=localhost port=5432 dbname=postgres user=myuser sslmode=disable"
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

