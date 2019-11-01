package message

import (
	"errors"
	"github.com/lishimeng/gouchen/internal/codec"
	"github.com/lishimeng/gouchen/internal/codec/intoyun"
	"github.com/lishimeng/gouchen/internal/codec/raw"
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
	switch encodeType {
	case codec.Javascript:
		// find from raw js repo
		data, err = raw.New().Encode(appId, msg)
		break
	case codec.IntoyunTLV:
		// find from tlv repo
		data, err = intoyun.New().Encode(appId, msg)
		break
	default:
		// no codec plugin
		err = errors.New("unknown codec type")
		break
	}
	return data, err
}
