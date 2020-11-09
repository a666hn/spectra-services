package main

import (
	"github.com/joho/godotenv"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"

	cl "github.com/skinnyguy/spectra-services/core/client"
	c "github.com/skinnyguy/spectra-services/core/proto"
	u "github.com/skinnyguy/spectra-services/core/utils"
	h "github.com/skinnyguy/spectra-services/microservices/handler"
)

func init() {
	if err := godotenv.Load(); err != nil {
		u.SendLogInfo("No .env file found...")
	}

	u.SendLogInfo("Data Source: ", u.GetDatasourceInfo())
}

func main() {
	var MicroVersion string = "latest"

	service := grpc.NewService(
		micro.Name(cl.OrganizationFamilyService),
		micro.Version(MicroVersion),
	)

	service.Init()
	c.RegisterOrganizationFamilyHandler(service.Server(), new(h.OrganizationFamilyHandler))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
