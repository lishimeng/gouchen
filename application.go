package gouchen

import (
	"fmt"
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/go-libs/shutdown"
	"github.com/lishimeng/gouchen/etc"
	"github.com/lishimeng/gouchen/internal/codec"
	"github.com/lishimeng/gouchen/internal/connector"
	"github.com/lishimeng/gouchen/internal/setup"
)

func Application(config etc.Configuration) (err error) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	log.Debug("start application")

	err = setup.Setup(config)
	if err != nil {
		return err
	}

	shutdown.WaitExit(&shutdown.Configuration{
		BeforeExit: func(s string) {
			log.Info("Shutdown [ %s ] (%s)", config.Name, s)
		},
	})

	return err
}

func Shutdown(msg string) {
	var message = "exit by user"
	if len(msg) > 0 {
		message = msg
	}
	shutdown.Exit(message)
}

func RegisterCodec(name string, b func() codec.Coder) {

	codec.RegisterBuilder(name, b)
}

func RegisterConnector(name string, b connector.Builder) {

	connector.RegisterBuilder(name, b)
}

func RegisterComponent(component func() error) {
	setup.CustomComponent(component)
}

func RegisterDbModel(model interface{}) {
	setup.RegisterDbModel(model)
}
