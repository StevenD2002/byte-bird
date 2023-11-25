package httpserver

import (
	"byte-bird/internal/service"
	"encoding/json"
	"net/http"
)

type HTTPServer struct {
	userService service.UserService
}

func NewHTTPServer(userService service.UserService) HTTPServer {
	return HTTPServer{userService}
}

func (s HTTPServer) StartServer() {
	http.HandleFunc("/register", s.handleRegisterUser)
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

