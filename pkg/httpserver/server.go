package httpserver

import (
	// "byte-bird/internal/service"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type HTTPServer struct {
}

func NewHTTPServer() HTTPServer {
  return HTTPServer{}
}


// Middleware for logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received:", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (s HTTPServer) StartServer() {
	http.Handle("/", loggingMiddleware(http.StripPrefix("/", http.FileServer(http.Dir("frontend")))))
	// serve the register page

  http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
}

