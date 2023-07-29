package main

import (
	"task/api"
	"task/cfg"
	"task/dbs"
	"task/internal/app/router"
	"task/pkg/db"
)

func main() {
	dbs.InitGPG(cfg.GetConfig())
	api.Serve(router.Routes(), cfg.GetAppPort())
	db.CloseDBConnection()
}
