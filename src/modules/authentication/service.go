package authentication

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Service interface {
	Registration(username string, password string) (err error, status int)
	Login(username string, password string) (err error, status int, token string)
}

type service struct {
	Db *sql.DB
}

func NewService(db *sql.DB) (svc *service) {
	svc = &service{
		Db: db,
	}

	return svc
}

func (s *service) Registration(username string, password string) (err error, status int) {
	status = fiber.StatusOK

	insertData := "INSERT INTO user (username, password) VALUES (?, ?)"

	_, err = s.Db.Exec(insertData, username, password)
	if err != nil {
		status = fiber.StatusInternalServerError
		return
	}

	return
}

func (s *service) Login(username string, password string) (err error, status int, token string) {
	status = fiber.StatusOK

	query := "SELECT * FROM user WHERE username = ? AND password = ?"

	userData := s.Db.QueryRow(query, username, password)

	var user User
	err = userData.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			status = fiber.StatusNotFound
			return
		}
		status = fiber.StatusInternalServerError
		return
	}

	token, err = generateToken(user.Username)
	if err != nil {
		status = fiber.StatusInternalServerError
		return
	}

	return
}

func generateToken(username string) (string, error) {
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("adjie.ganteng.banget"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
