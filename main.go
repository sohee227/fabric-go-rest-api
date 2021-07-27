package main

import (
	"fabric-go-rest-api/server"
)

func main() {
	s := server.NewServer()

	if err := s.Init(3000); err != nil {
		panic(err)
	}

	s.Start()
}
