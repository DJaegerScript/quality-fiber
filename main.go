package main

import (
	"djaeger-software-testing/src/config"
	"log"
)

func main() {
	s := config.MakeServer()
	defer func() {
		s.DB.Close()
	}()

	err := s.Router.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
