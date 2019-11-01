package setup

import (
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/gouchen/etc"
	"github.com/lishimeng/gouchen/internal/downlink"
	"time"
)

func setupDownLink() error {
	log.Debug("setup downLink")
	downlink.Init(etc.Config.DownLink.FetchSize, time.Duration(etc.Config.DownLink.IdleTime)*time.Millisecond)
	run()
	return nil
}

func run() {
	go func() {
		handler := downlink.GetInstance()
		h := *handler
		h.StartDownLink()
	}()
}
