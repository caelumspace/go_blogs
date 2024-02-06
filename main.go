package main

import (
	"fmt"

	"capture-life-api/app"
	"capture-life-api/config"
)

func main() {
	MongoURI := config.MongoURI
	fmt.Println("Starting application...")
	app := &app.App{}
	app.Initialize(MongoURI)
	var port string = "5000"
	fmt.Println(`Server running @` + port)
	app.Run(":" + port)
}
