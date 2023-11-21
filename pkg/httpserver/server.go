
package httpserver

import (
	"encoding/json"
	"net/http"

	"byte-bird/internal/service"
  "byte-bird/internal/domain/post"
)

type HTTPServer struct {
	userService service.UserService
  postService service.PostService
}

func NewHTTPServer(userService service.UserService, postService service.PostService) HTTPServer {
	return HTTPServer{userService, postService}
}

func (s HTTPServer) StartServer() {
	http.HandleFunc("/create-user", s.handleCreateUser)
  http.HandleFunc("/create-post", s.handleCreatePost)
	http.ListenAndServe(":8080", nil)
}

func (s HTTPServer) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.userService.CreateUser(user.Name, user.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s HTTPServer) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newPost struct {
		Content string `json:"content"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newPost); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a post using the PostService
	post := &post.Post{
		Content: newPost.Content,
		// Add other relevant fields
	}

	if err := s.postService.CreatePost(r.Context(), post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
