package sqlstore_test

import (
	"testing"
	"github.com/kantegory/http-rest-api/internal/app/store/sqlstore"
	"github.com/kantegory/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
)


// TestUserRepositury_Create ...
func TestUserRepositury_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown("users")

	s := sqlstore.New(db)
	u, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

// TestUserRepositury_FindByEmail ...
func TestUserRepositury_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email

	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
