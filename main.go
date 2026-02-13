package main

import (
	"fmt"
	"log"

	"github.com/polishedfeedback/aniverse/services"
)

func main() {
	anime, err := services.SearchAnime("jujutsu kaisen")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(anime[0])
}
