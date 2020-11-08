package main

import (
	"github.com/joho/godotenv"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"

	coreclient "github.com/skinnyguy/spectra-services/core/client"
	core "github.com/skinnyguy/spectra-services/core/proto"
	utils "github.com/skinnyguy/spectra-services/core/utils"
	handler "github.com/skinnyguy/spectra-services/microservices/handler"
)

const MicroVersion string = "latest"

func init() {
	if err := godotenv.Load(); err != nil {
		utils.SendLogInfo("No .env file found...")
	}

	utils.SendLogInfo("[Data Source] : ", utils.GetDatasourceInfo())
}

func main() {
	service := grpc.NewService(
		micro.Name(coreclient.AccountService),
		micro.Version(MicroVersion),
	)

	service.Init()
	core.RegisterAccountServicesHandler(service.Server(), new(handler.AccountHandler))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
