package setup

import (
	"github.com/lishimeng/go-libs/web"
	"github.com/lishimeng/gouchen/internal/api"
	"github.com/lishimeng/gouchen/internal/webserver"
)

func setupWebServer() error {

	var components = []web.Component{
		api.SetupDataPoint,
		api.SetupApplication,
		api.SetupLogic,
		api.SetupConnector,
		api.SetupCodecJs,
		api.SetupTrigger,
	}
	go webserver.Run(components...)
	return nil
}
