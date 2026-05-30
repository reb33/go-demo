package main

import (
	"fmt"
	"net/http"

	"adv_demo/internal/hello"
	"adv_demo/internal/auth"
)

func main() {
	// conf := configs.LoadConfig()
	router := http.NewServeMux()
	hello.NewHelloHandler(router)
	auth.NewAuthHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
