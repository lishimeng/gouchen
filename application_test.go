package gouchen

import (
	"github.com/lishimeng/gouchen/internal/codec"
	"testing"
)

type sampleBuilder struct {

}

func (s *sampleBuilder) Decode(appId string, data []byte) (props map[string]interface{}, err error) {
	return props, err
}

func (s *sampleBuilder) Encode(appId string, props map[string]interface{}) (data []byte, err error) {
	data = []byte{0x01}
	return data, err
}


func TestRegisterCodec(t *testing.T) {
	var name = "test_codec"
	var builder = func() codec.Coder {
		tmp := sampleBuilder{}
		var h codec.Coder = &tmp
		return h
	}

	RegisterCodec(name, builder)
	b, err := codec.GetBuilder(name)
	if err != nil {
		t.Fatal(err)
		return
	}
	res, err := b().Encode("appId", make(map[string]interface{}))
	if err != nil {
		t.Fatal(err)
		return
	}
	if len(res) != 1 {
		t.Fatalf("expect data length 1, but %d", len(res))
		return
	}

	if res[0] != 0x01 {
		t.Fatalf("expect data length 0x01, but %b", res[0])
		return
	}
}
