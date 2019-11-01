package setup

import "github.com/lishimeng/gouchen/internal/integration/persistent"

func setupInflux() error {
	err := persistent.Init()
	return err
}
