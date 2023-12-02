package authentication

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-faker/faker/v4"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

type AuthenticationData struct {
	Username string `faker:"username"`
	Password string `faker:"password"`
}

func TestRegistrationHandler(t *testing.T) {
	db, err := sql.Open("sqlite3", "../../../softtest.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	authenticationHandler := NewHandler(db)
    
	app := fiber.New()

	app.Post("/auth/registration", authenticationHandler.Register)

	var registrationData AuthenticationData
	err = faker.FakeData(&registrationData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	body := strings.NewReader(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, registrationData.Username, registrationData.Password))
	req := httptest.NewRequest(http.MethodPost, "/auth/registration", body)
	req.Header.Set("Content-Type", "application/json")
	
	resp, _ := app.Test(req)
	
	assert.Equal(t, 201, resp.StatusCode)
}

func TestLoginHandler(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM user WHERE username = ? AND password = ?")).
		WithArgs("test_user", "test_password").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password"}).AddRow(1, "test_user", "test_password"))

	authenticationHandler := NewHandler(db)
    
	app := fiber.New()

	app.Post("/auth/login", authenticationHandler.Login)

	body := strings.NewReader(fmt.Sprintf(`{"username": "test_user", "password": "test_password"}`))
	req := httptest.NewRequest(http.MethodPost, "/auth/login", body)
	req.Header.Set("Content-Type", "application/json")
	
	resp, _ := app.Test(req)
	
	assert.Equal(t, 200, resp.StatusCode)
}