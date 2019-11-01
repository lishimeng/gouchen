package message

import (
	"github.com/lishimeng/gouchen/internal/codec"
)

func decode(appId string, decodeType string, rawData []byte) (data map[string]interface{}, err error) {

	coder, err := codec.GetBuilder(decodeType)
	if err != nil {
		return data, err
	}
	data, err = coder().Decode(appId, rawData)
	return data, err
}

func encode(appId string, encodeType string, msg map[string]interface{}) (data []byte, err error) {
	coder, err := codec.GetBuilder(encodeType)
	if err != nil {
		return data, err
	}
	data, err = coder().Encode(appId, msg)
	return data, err
}
