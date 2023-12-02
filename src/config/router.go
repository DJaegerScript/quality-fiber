package config

import (
	"djaeger-software-testing/src/modules/authentication"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (s *Server) SetupRouter() {
	s.Router.Use(recover.New())

	s.Router.Mount("/auth", s.AuthenticationRouter())
}

func (s *Server) AuthenticationRouter() *fiber.App {
	authenticationHandler := authentication.NewHandler(s.DB)

	authRouter := fiber.New()

	authRouter.Post("/registration", authenticationHandler.Register)
	authRouter.Post("/login", authenticationHandler.Login)

	return authRouter
}
