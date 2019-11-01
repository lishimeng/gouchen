package connector

import (
	"fmt"
	"github.com/lishimeng/gouchen/internal/model"
)

const (
	// lorawan
	LoraWanType = "lora"
	MqttJson    = "mqtt_json"
	Amqp        = "amq"
)

type Config struct {
	ID    string            `json:"id"`
	Name  string            `json:"name"`
	Type  string            `json:"type"`
	Props map[string]string `json:"props"`
}

type UpLinkListener func(context *model.LinkMessage)
type Builder func(conf Info, props map[string]string, listener UpLinkListener) (c *Connector, err error)

type Connector interface {
	GetID() string
	GetName() string
	GetType() string
	GetState() bool
	// 即时发送
	DownLink(target model.Target, data []byte)
}

var builders = make(map[string]Builder)

func RegisterBuilder(name string, builder Builder) {

	builders[name] = builder
}

func GetBuilder(name string) (b Builder, err error) {
	if item, ok := builders[name]; ok {
		b = item
	} else {
		err = fmt.Errorf("unknown connector type %s", name)
	}
	return b, err
}
