package main

import (
	"fmt"
	"greenpipe/internal/server"
)

func main() {

	server := server.NewServer()

	fmt.Println("Starting server on port 8080")
	
	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
