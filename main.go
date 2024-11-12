package main

import (
	"be-awards/config"
	"be-awards/route"
)

func main() {
	config.InitDB()

	r := route.InitRoutes()
	r.Run(":8000")
}
