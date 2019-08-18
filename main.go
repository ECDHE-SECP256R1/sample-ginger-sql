package main

import (
	"github.com/kulichak/sample-ginger-sql/project/controllers"
)

func main() {
	engine := controllers.RegisterRoutes()
	engine.Run()
}
