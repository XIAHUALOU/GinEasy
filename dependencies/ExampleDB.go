package dependencies

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

type XOrmAdapter struct {
	*xorm.Engine
}

func NewXOrmAdapter() *XOrmAdapter {
	engine, err := xorm.NewEngine("mysql",
		"root:123456@tcp(xxx.xxx.xxx.xxx:xxxx)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	engine.DB().SetMaxIdleConns(5)
	engine.DB().SetMaxOpenConns(10)
	return &XOrmAdapter{Engine: engine}
}
