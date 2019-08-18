package v1

import (
	g "github.com/kulichak/ginger"
	"github.com/kulichak/models"
)

type HelloItemsController struct {
	g.BaseItemsController
}

func (c *HelloItemsController) Post(request *models.Request) {
	c.BaseItemsController.Post(request)
}

func (c *HelloItemsController) Get(request *models.Request) {
	c.BaseItemsController.Get(request)
}

type HelloItemController struct {
	g.BaseItemController
}

func (c *HelloItemController) Get(request *models.Request) {
	c.BaseItemController.Get(request)
}
