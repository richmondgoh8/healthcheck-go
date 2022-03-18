package main

import (
	"log"

	"github.com/rlc4u/health-check/internal/router"
)

func main() {

	r := router.SetupRouter()
	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatal(err, "unable to start server")
	}
}
