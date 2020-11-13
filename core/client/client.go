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
	// Role service ...
	PrivilegeService = GoMicroPrefix + "privileges.service"
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
	privilegeOnce    sync.Once
	privilegeService service.PrivilegeService
)

// GetRoleService ...
func GetPrivilegeService() service.PrivilegeService {
	privilegeOnce.Do(func() {
		options := getClientOptions()
		privilegeService = service.NewPrivilegeService(PrivilegeService, grpcclient.NewClient(options...))
	})

	return privilegeService
}
