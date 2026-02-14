package main

import (
	"log"

	"github.com/polishedfeedback/aniverse/router"
)

func main() {
	r := router.SetupRouter()

	err := r.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
