package main

import (
	"restful-api-practice/code-structuring/config"
	"restful-api-practice/code-structuring/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}