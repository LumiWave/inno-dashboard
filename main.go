package main

import (
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/inno-dashboard/rest_server/app"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	if err := app.Start(); err != nil {
		log.Errorf("%v", err)
		return
	}
}
