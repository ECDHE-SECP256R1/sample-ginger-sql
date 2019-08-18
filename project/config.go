package project

import (
	"github.com/kulichak/models"
	"github.com/kulichak/sample-ginger-sql/project/helpers"
	"github.com/kulichak/sql"
)

type Config struct {
	models.IConfig

	Port string

	// sql config
	SqlDialect          string
	SqlConnectionString string
	//
}

var CurrentConfig *Config

func init() {
	CurrentConfig = &Config{
		Port:       helpers.GetEnv("Port", "8080"),
		SqlDialect: helpers.GetEnv("SqlDialect", "mysql"),
		SqlConnectionString: helpers.GetEnv("SqlConnectionString",
			"root:123@/test?charset=utf8&parseTime=True&loc=UTC"),
	}
	sql.InitializeConfig(CurrentConfig)
}
