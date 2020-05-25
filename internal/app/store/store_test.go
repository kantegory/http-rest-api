package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

// TestMain ...
func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "host=localhost dbname=restapi_test sslmode=disable user=YOURUSER password=YOURPASSWORD"
	}

	os.Exit(m.Run())
}