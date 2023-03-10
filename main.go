package main

import (
	"fmt"
	"go-identity-manager/cmd/router"
	"log"
	"net/http"
)

func main() {
	routes := router.CreateRouter()

	fmt.Printf("Server running on port %d\n", 9696)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 9696), routes))
}
