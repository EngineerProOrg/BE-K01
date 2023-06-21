package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/EngineerProOrg/BE-K01/configs"
	_ "github.com/EngineerProOrg/BE-K01/docs"
	"github.com/EngineerProOrg/BE-K01/internal/app/web_app"
	"github.com/EngineerProOrg/BE-K01/internal/app/web_app/service"
)

var (
	path = flag.String("config", "config.yml", "config path for this service")
)

// @title           Gin Social network Service
// @version         1.0
// @description     A simple social network management service API in Go using Gin framework.
// @termsOfService

// @contact.name   Dong Truong
// @contact.url    https://www.linkedin.com/in/dong-truong-56297a145/
// @contact.email  tpdongcs@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath  /api/v1
//	@securitydefinitions.oauth2.password	OAuth2Password
//	@tokenUrl								https://example.com/oauth/token
//	@scope.read								Grants read access
//	@scope.write							Grants write access
//	@scope.admin							Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode	OAuth2AccessCode
// @tokenUrl								https://example.com/oauth/token
// @authorizationUrl						https://example.com/oauth/authorize
// @scope.admin							Grants read and write access to administrative information
func main() {
	flag.Parse()
	conf, err := configs.GetWebConfig(*path)
	fmt.Println(conf)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
	webSvc, err := service.NewWebService(conf)
	if err != nil {
		log.Fatalf("failed to init service: %v", err)
	}
	web_app.WebController{
		WebService: *webSvc,
		Port:       conf.Port,
	}.Run()
}
