package main

import (
	."github.com/erdoganakdeniz/app"
	."github.com/erdoganakdeniz/config"
	."github.com/erdoganakdeniz/router"
	"log"
)

func main() {
	app:=SetupApp()
	SetupDB()
	SetupRouter(app)
	if err:=app.Listen(8080);err != nil {
		log.Fatal(err)
	}
}
