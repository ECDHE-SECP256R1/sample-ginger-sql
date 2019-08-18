package logics

import (
	"github.com/kulichak/logic"
	"github.com/kulichak/models"
	m "github.com/kulichak/sample-ginger-sql/project/models"
)

type HelloLogic struct {
	logic.BaseLogicHandler
}

func (hello *HelloLogic) Model(request *models.Request) {
	request.Model = &m.Hello{}
}

func (hello *HelloLogic) Models(request *models.Request) {
	request.Models = &[]m.Hello{}
}

func (hello *HelloLogic) BeforeQuery(request *models.Request) {
	request.Filters = &models.Filters{

	}
}
