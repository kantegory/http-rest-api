package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/kantegory/http-rest-api/internal/app/store/teststore"
)

// TestServer_HandleUsersCreate ...
func TestAPIServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New())

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)

	s.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusOK)
}
