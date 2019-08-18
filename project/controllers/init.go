package controllers

import (
	g "github.com/kulichak/ginger"
	"github.com/kulichak/sample-ginger-sql/project"
	v1 "github.com/kulichak/sample-ginger-sql/project/controllers/v1"
)

func RegisterRoutes() *g.Router {
	router := g.NewRouter()
	router.Address = []string{"0.0.0.0:" + project.CurrentConfig.Port}
	routerV1 := router.Group("/v1")
	v1.RegisterRoutes(routerV1)
	return router
}
