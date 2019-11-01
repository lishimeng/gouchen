package db

import (
	"github.com/astaxie/beego/orm"
	"github.com/lishimeng/go-libs/persistence"
)

var Orm persistence.OrmContext

func RegisterDriver(driverName string, typ int) error {
	return orm.RegisterDriver(driverName, orm.DriverType(typ))
}

func Init(config persistence.PostgresConfig) (err error) {

	Orm, err = persistence.InitPostgresOrm(config)
	return err
}

func RegisterModel(model interface{}) {
	orm.RegisterModel(model)
}
