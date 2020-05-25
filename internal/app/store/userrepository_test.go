package store_test

import (
	"testing"
	"github.com/kantegory/http-rest-api/internal/app/store"
	"github.com/kantegory/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
)


// TestUserRepositury_Create ...
func TestUserRepositury_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)

	defer teardown("users")

	u, err := s.User().Create(&model.User {
		Email: "user@example.org",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
