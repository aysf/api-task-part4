package main

import (
	"aysf/day6r1/config"
	"aysf/day6r1/routes"

	m "aysf/day6r1/middlewares"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8080"))
}
