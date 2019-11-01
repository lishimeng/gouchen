package codec

import "fmt"

const (
	None       = "none"
	Javascript = "raw"
	Protobuf   = "protobuf"
	IntoyunTLV = "intoyuntlv"
)

type Coder interface {
	Decode(appId string, data []byte) (props map[string]interface{}, err error)
	Encode(appId string, props map[string]interface{}) (data []byte, err error)
}

var builders =make(map[string]func() Coder)

func RegisterBuilder(name string, builder func() Coder) {

	builders[name] = builder
}

func GetBuilder(name string) (b func() Coder, err error) {
	if item, ok := builders[name]; ok {
		b = item
	} else {
		err = fmt.Errorf("unknown codec type %s", name)
	}
	return b, err
}