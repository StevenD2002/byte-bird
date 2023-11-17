// internal/pkg/httpserver/server_test.go

package httpserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"byte-bird/internal/service"

	"github.com/stretchr/testify/assert"

)

func TestCreateUserEndpoint(t *testing.T) {
	// Initialize the mock UserService for testing
	mockUserService := service.NewMockUserService()

	// Create a new HTTPServer instance for testing, using the mock UserService
	server := NewHTTPServer(mockUserService)

	// Set up a test server
	ts := httptest.NewServer(http.HandlerFunc(server.handleCreateUser))
	defer ts.Close()

	// Create a sample user payload
	userPayload := map[string]string{"name": "John Doe", "email": "john.doe@example.com"}
	jsonPayload, _ := json.Marshal(userPayload)

	// Send a POST request to the test server
	resp, err := http.Post(ts.URL+"/create-user", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Assert the response status code
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
