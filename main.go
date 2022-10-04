package main

import (
	"fmt"
	"log"

	"github.com/depri11/be_e-commerce/app/router"
	"github.com/depri11/be_e-commerce/common/configs"
	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	configs := configs.NewConfiguration()
	config, err := configs.Read()
	if err != nil {
		log.Println(err)
		return
	}

	urlServe := fmt.Sprintf(":%d", config.PortApi)
	log.Println("running at", urlServe)

	app := router.Setup(config)
	log.Fatal(app.Listen(urlServe))
}
