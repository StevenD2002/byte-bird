package httpserver

import (
    "encoding/json"
    "net/http"

    "github.com/stevend2002/tgp-bp/internal/service"
)

type HTTPServer struct {
    userService service.UserService
}

func NewHTTPServer(userService service.UserService) HTTPServer {
    return HTTPServer{userService}
}

func (s HTTPServer) StartServer() {
    http.HandleFunc("/create-user", s.handleCreateUser)
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

