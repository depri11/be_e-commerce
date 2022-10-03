package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

}
