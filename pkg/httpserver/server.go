package httpserver

import (
	"byte-bird/internal/domain/post"
	"byte-bird/internal/service"
	"context"
	"encoding/json"
	"net/http"
  "time"
)

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
  http.HandleFunc("/createPost", s.handleCreatePost)
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
	ctx := context.Background()
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	userID := "1234" // for now we will hardcode this
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
		UserID:  userID,
		Content: newPost.Content,
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
