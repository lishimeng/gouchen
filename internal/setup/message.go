package setup

import (
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/gouchen/internal/message"
	"github.com/lishimeng/gouchen/internal/model"
)

func setupMessage() (err error) {
	message.GetEngine().SetCallback(onDownLink)
	return err
}

func onDownLink(target model.Target, data []byte) {
	conn, err := ConnectorRepository.GetByID(target.ConnectorId)
	if err != nil {
		log.Debug("no connector, skip this data application:%s connector:%s", target.AppId, target.ConnectorId)
	} else {
		c := *conn
		log.Debug("down link to %s, %s, %s", c.GetID(), c.GetName(), c.GetType())
		c.DownLink(target, data)
		log.Debug("down link completed %s:%s:%s", target.ConnectorId, target.AppId, target.DeviceId)
	}
}
