package httpserver

import (
	"byte-bird/internal/domain/post"
	"byte-bird/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type HTTPServer struct {
	userService service.UserService
	postService service.PostService
}

func NewHTTPServer(userService service.UserService, postService service.PostService) HTTPServer {
	return HTTPServer{userService, postService}
}

func (s HTTPServer) StartServer() {
	http.HandleFunc("/register", s.handleRegisterUser)
	http.HandleFunc("/login", s.handleLoginUser)
	http.HandleFunc("/createPost", AuthenticateMiddleware(s.handleCreatePost))
	http.ListenAndServe(":8080", nil)
}

func (s HTTPServer) handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Start the create user process
	err := s.userService.CreateUser(user.Name, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s HTTPServer) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	// Start the login user process
	token, err := s.userService.AuthenticateUser(ctx, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Error logging in user", http.StatusInternalServerError)
		return
	}
	// build the response object

	response := map[string]interface{}{
		"status": "success",
		"token":  token,
	}

	jsonRespone, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRespone)
}

// create post with the body and users name
func (s HTTPServer) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // Retrieve the context from the request

	// Extract UserID from the context
	userID, ok := ctx.Value(userIDKey).(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newPost struct {
		Content string `json:"content"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newPost); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// create the new post
	post := &post.Post{
		UserID:    userID, // Use the extracted UserID
		Content:   newPost.Content,
		Timestamp: time.Now(),
	}

	// start the create post process
	err := s.postService.CreatePost(ctx, post)
	if err != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Claims structure to represent the JWT claims
// Key used to store user ID in the context
const userIDKey = "user_id"

// AuthenticateMiddleware is a middleware for extracting and validating JWT tokens
func AuthenticateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Remove the "Bearer " prefix
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Parse the token with custom claims
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("temp-secret-key"), nil // Replace with your actual secret key
		})

		if err != nil || !token.Valid {
			fmt.Println("Token validation failed:", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extract claims from the token
		claims, ok := token.Claims.(*Claims)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Populate the context with the user ID from the token
		ctx := context.WithValue(r.Context(), userIDKey, claims.UserID)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	}
}
