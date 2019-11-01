package message

import (
	"github.com/lishimeng/gouchen/internal/model"
)

type OnDownLink func(target model.Target, data []byte)

type DataProcessEngine struct {
	cb OnDownLink
}

var singleton *DataProcessEngine

func init() {
	d := DataProcessEngine{}
	singleton = &d
}

func GetEngine() *DataProcessEngine {
	return singleton
}

func (d *DataProcessEngine) SetCallback(cb OnDownLink) {
	d.cb = cb
}

func (d DataProcessEngine) OnDataUpLink(upLink *model.LinkMessage) {
	processUpLink(upLink)
}

func (d DataProcessEngine) OnDataDownLink(target model.Target, props map[string]interface{}) {
	processDownLink(target, props, d.cb)
}
