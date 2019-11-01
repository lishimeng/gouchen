package setup

import (
	"encoding/json"
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/gouchen/internal/connector"
	"github.com/lishimeng/gouchen/internal/connector/connectors/lorawan"
	"github.com/lishimeng/gouchen/internal/connector/connectors/mq"
	"github.com/lishimeng/gouchen/internal/connector/connrepo"
	"github.com/lishimeng/gouchen/internal/db/repo"
	"github.com/lishimeng/gouchen/internal/message"
	"time"
)

var ConnectorRepository *connrepo.Repository

func setupConnector() error {

	connector.RegisterBuilder(connector.LoraWanType, lorawan.Create)
	connector.RegisterBuilder(connector.MqttJson, mq.Create)

	ConnectorRepository = connrepo.New()

	loadConnectors()
	return nil
}

func loadConnectors() {

	var _initAllConnector = func() {

		log.Debug("check connectors")
		// get all config
		configs, size := repo.GetConnectConfigs()
		if size > 0 {
			// loop config to connect message platform
			log.Debug("connector size:%d", size)
			for _, config := range configs {
				loadConnector(*config)
			}
		}
	}

	for {
		_initAllConnector()
		time.Sleep(time.Second * 10)
		break
	}
}

func loadConnector(connConf repo.ConnectorConfig) {

	if !connExist(connConf.Name) {
		log.Debug("load connector[%s]", connConf.Name)
		var props map[string]string
		err := json.Unmarshal([]byte(connConf.Props), &props)
		if err != nil {
			return
		}

		config := connector.Config{
			ID:    connConf.Id,
			Name:  connConf.Name,
			Type:  connConf.Type,
			Props: props,
		}
		var c *connector.Connector
		c, err = createConn(config)
		if err != nil {
			return
		}
		ConnectorRepository.Register(c)
	}
}

func createConn(conf connector.Config) (c *connector.Connector, err error) {
	var t = conf.Type
	var builder connector.Builder
	builder, err = connector.GetBuilder(t)
	if err == nil {
		messageEngine := message.GetEngine()
		info := connector.Info{
			Id:   conf.ID,
			Name: conf.Name,
			Type: t,
		}
		c, err = builder(info, conf.Props, messageEngine.OnDataUpLink)
	}
	return c, err
}

func connExist(id string) bool {
	_, err := ConnectorRepository.GetByID(id)
	// TODO check status
	return err == nil
}
