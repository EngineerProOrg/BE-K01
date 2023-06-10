package main

import (
	"flag"
	"log"

	"github.com/EngineerProOrg/BE-K01/configs"
	"github.com/EngineerProOrg/BE-K01/internal/app/web_app"
	"github.com/EngineerProOrg/BE-K01/internal/app/web_app/service"
)

var (
	path = flag.String("conf", "config.yml", "config path for this service")
)

func main() {
	conf, err := configs.GetWebConfig(*path)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
	webSvc, err := service.NewWebService(conf)
	if err != nil {
		log.Fatalf("failed to init service: %v", err)
	}
	web_app.WebController{
		WebService: *webSvc,
	}.Run()
}
