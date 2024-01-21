package main

import (
	// "database/sql"
	// "fmt"
	// "byte-bird/internal/repository"
	// "byte-bird/internal/service"

	// "net/http"

	"fmt"
	"net/http"

	"github.com/a-h/templ"
	_ "github.com/lib/pq"
	// "byte-bird/pkg/errors"
)

const (
	// will change this later and throw stuff in an env
	dbConnectionString = "host=localhost port=5432 dbname=mydatabase user=postgres password=password sslmode=disable"
)

func main() {
	component := hello("John")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
