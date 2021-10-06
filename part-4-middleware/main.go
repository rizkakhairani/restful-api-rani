package main

import (
	"restful-api-practice/middleware/config"
	"restful-api-practice/middleware/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}