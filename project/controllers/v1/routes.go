package v1

import (
	g "github.com/kulichak/ginger"
	"github.com/kulichak/sample-ginger-sql/project/db"
	"github.com/kulichak/sample-ginger-sql/project/logics"
)

var hello = new(HelloItemsController)
var helloItem = new(HelloItemController)

func init() {
	hello.BaseController.Init(&logics.HelloLogicHandler, &db.HelloDbHandler)
	helloItem.BaseController.Init(&logics.HelloLogicHandler, &db.HelloDbHandler)
}

func RegisterRoutes(router *g.RouterGroup) {
	hello.AddRoute("Get")
	hello.AddRoute("Post")
	helloItem.AddRoute("Get")

	hello.RegisterRoutes(hello, "/hello", router)
	helloItem.RegisterRoutes(helloItem, "/hello/:id", router)
}
