package setup

import (
	_ "github.com/lib/pq"
	"github.com/lishimeng/go-libs/persistence"
	"github.com/lishimeng/gouchen/internal/db"
	"github.com/lishimeng/gouchen/internal/db/repo"
	"github.com/lishimeng/gouchen/internal/etc"
)

var databaseModels []interface{}

func defaultDbModels() {
	//RegisterDbModel(new(repo.DeviceConfig))
	RegisterDbModel(new(repo.AppConfig))
	RegisterDbModel(new(repo.DataPoint))
	RegisterDbModel(new(repo.LogicScript))
	RegisterDbModel(new(repo.ConnectorConfig))
	RegisterDbModel(new(repo.CodecScript))
	RegisterDbModel(new(repo.DownLinkData))
	RegisterDbModel(new(repo.DelayedDownLinkData))
	RegisterDbModel(new(repo.DownLinkLog))
	RegisterDbModel(new(repo.TriggerConfig))
}

func dbRepo() (err error) {
	var config = persistence.PostgresConfig{
		UserName: etc.Config.Db.User,
		Password: etc.Config.Db.Password,
		Host:     etc.Config.Db.Host,
		Port:     etc.Config.Db.Port,
		DbName:   etc.Config.Db.Database,
		MaxIdle:  5,  // TODO move into config file
		MaxConn:  10, // TODO move into config file
		InitDb:   true,
	}

	defaultDbModels()

	err = db.RegisterDriver("postgres", 4)
	if err != nil {
		return err
	}
	if len(databaseModels) > 0 {
		for model := range databaseModels {
			db.RegisterModel(model)
		}
	}
	err = db.Init(config)

	return err
}

func RegisterDbModel(model interface{}) {
	databaseModels = append(databaseModels, model)
}
