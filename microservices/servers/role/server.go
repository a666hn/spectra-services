package main

import (
	"github.com/joho/godotenv"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"

	cc "github.com/skinnyguy/spectra-services/core/client"
	cp "github.com/skinnyguy/spectra-services/core/proto"
	ut "github.com/skinnyguy/spectra-services/core/utils"
	h "github.com/skinnyguy/spectra-services/microservices/handler"
)

func init() {
	if err := godotenv.Load(); err != nil {
		ut.SendLogInfo("No .env file found...")
	}

	ut.SendLogInfo("Datasource: ", ut.GetDatasourceInfo())
}

func main() {
	var MicroVersion string = "latest"

	service := grpc.NewService(
		micro.Name(cc.RoleService),
		micro.Version(MicroVersion),
	)

	service.Init()
	cp.RegisterRoleHandler(service.Server(), new(h.RoleHandler))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
