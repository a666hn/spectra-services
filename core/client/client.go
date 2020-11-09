package client

import (
	"sync"
	"time"

	configFile "github.com/skinnyguy/spectra-services/core/config"
	service "github.com/skinnyguy/spectra-services/core/proto"

	"github.com/micro/go-micro/client"
	grpcclient "github.com/micro/go-micro/client/grpc"
)

const GoMicroPrefix string = "go.micro.srv."

const (
	// Account service ...
	AccountService = GoMicroPrefix + "account.service"
	// Organization Family service ...
	OrganizationFamilyService = GoMicroPrefix + "organization-family.service"
)

func getClientOptions() []client.Option {
	config := configFile.GetSpectraConfig()
	options := []client.Option{
		client.RequestTimeout(config.GRPCTimeout * time.Second),
	}

	return options
}

/* ========================== */
// Define all client services //
/* ========================== */

var (
	accountOnce    sync.Once
	accountService service.AccountServicesService
)

// GetAccountService ...
func GetAccountService() service.AccountServicesService {
	accountOnce.Do(func() {
		options := getClientOptions()
		accountService = service.NewAccountServicesService(AccountService, grpcclient.NewClient(options...))
	})

	return accountService
}

var (
	orgOne     sync.Once
	orgService service.OrganizationFamilyService
)

// GetOrganizationFamilyService ...
func GetOrganizationFamilyService() service.OrganizationFamilyService {
	orgOne.Do(func() {
		options := getClientOptions()
		orgService = service.NewOrganizationFamilyService(OrganizationFamilyService, grpcclient.NewClient(options...))
	})

	return orgService
}
