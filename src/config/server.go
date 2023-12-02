package config

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Router *fiber.App
	DB     *sql.DB
}

func MakeServer() Server {
	r := fiber.New()
	db := InitDatabase()
	server := Server{
		Router: r,
		DB:     db,
	}

	server.SetupRouter()

	return server
}
