package setup

import (
	"github.com/lishimeng/gouchen/internal/event"
	"github.com/lishimeng/gouchen/internal/message"
	"github.com/lishimeng/gouchen/internal/model"
)

func setupEvent() (err error) {
	event.GetInstance().SetCallback(onEvent)
	return err
}

func onEvent(target model.Target, properties map[string]interface{}) {
	message.GetEngine().OnDataDownLink(target, properties)
}
