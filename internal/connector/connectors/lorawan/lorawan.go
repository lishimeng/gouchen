package lorawan

import (
	"encoding/base64"
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/go-connector/lorawan"
	"github.com/lishimeng/gouchen/internal/connector"
	"github.com/lishimeng/gouchen/internal/model"
)

type connectorLoraWan struct {
	connector.Info
	Proxy    *lorawan.Connector
	Listener connector.UpLinkListener
}

// TODO qos
func New(info connector.Info,
	broker string,
	clientId string,
	topicUpLink string,
	topicDownLink string,
	listener connector.UpLinkListener) (connector.Connector, error) {

	log.Debug("Lorawan connector[%s]", broker)

	c := connectorLoraWan{
		Info:     info,
		Listener: listener,
	}

	proxy, err := lorawan.New(broker, clientId, topicUpLink, topicDownLink, 0)
	if err != nil {
		return nil, err
	}
	c.Proxy = proxy
	proxy.SetUpLinkListener(c.onMessage)
	err = proxy.ConnectOnce()
	if err != nil {
		log.Debug(err)
	}

	var conn connector.Connector = &c
	return conn, nil // TODO
}

func Create(conf connector.Info, props map[string]string, listener connector.UpLinkListener) (c *connector.Connector, err error) {

	var con connector.Connector
	con, err = New(
		conf,
		props["broker"],
		props["clientId"],
		props["upLink"],
		props["downLink"],
		listener,
	)
	if err == nil {
		c = &con
	}
	return c, err
}

func (c connectorLoraWan) GetState() bool {
	return c.Proxy.GetSession().State
}

// 监听数据上传
///
func (c *connectorLoraWan) onMessage(payload lorawan.PayloadRx) {
	rawData, err := base64.StdEncoding.DecodeString(payload.Data)
	if err != nil {
		return
	}
	context := model.LinkMessage{}
	context.ApplicationID = payload.ApplicationID
	context.DeviceID = payload.DevEUI
	context.ApplicationName = payload.ApplicationName
	context.DeviceName = payload.DeviceName
	context.Raw = rawData

	// 解析object字段
	if payload.DataObj != nil {
		context.Data = *payload.DataObj
	}

	c.Listener(&context)
}

func (c connectorLoraWan) DownLink(target model.Target, logicData []byte) {
	// 业务数据部分必须为base64格式
	raw := base64.StdEncoding.EncodeToString(logicData)
	log.Debug("lora down link [%s]%s:%s", target.ConnectorId, target.AppId, target.DeviceId)
	log.Debug("data object:%s", logicData)
	log.Debug("data raw:%s", raw)

	downLinkData := lorawan.PayloadTx{FPort: 3, Data: raw}

	go func() {
		err := c.Proxy.DownLink(target.AppId, target.DeviceId, downLinkData)
		if err != nil {
			log.Debug(err)
		}
	}()
}
