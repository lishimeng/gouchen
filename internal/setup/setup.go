package setup

import "github.com/lishimeng/gouchen/internal/etc"

var components = []func() error{
	dbRepo,
	setupCodec,
	setupEvent,
	setupMessage,
	setupWebServer,
	setupInflux,
	setupDownLink,
	setupConnector,
}

func Setup(config etc.Configuration) (err error) {

	etc.Config = &config

	components := []func() error{
		dbRepo,
		setupEvent,
		setupMessage,
		setupWebServer,
		setupInflux,
		setupDownLink,
		setupConnector,
	}

	for _, component := range components {
		if err := component(); err != nil {
			return err
		}
	}
	return err
}

func CustomComponent(component func() error) {
	components = append(components, component)
}
