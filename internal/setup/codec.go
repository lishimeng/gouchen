package setup

import (
	"github.com/lishimeng/gouchen/internal/codec"
	"github.com/lishimeng/gouchen/internal/codec/intoyun"
	"github.com/lishimeng/gouchen/internal/codec/raw"
)

func setupCodec() error {

	codec.RegisterBuilder(codec.IntoyunTLV, intoyun.New)
	codec.RegisterBuilder(codec.Javascript, raw.New)
	return nil
}
