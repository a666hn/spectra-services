module github.com/skinnyguy/spectra-services/microservices

go 1.14

require (
	github.com/joho/godotenv v1.3.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/kubernetes v0.0.0-20200119172437-4fe21aa238fd
	github.com/skinnyguy/spectra-services/core v0.0.0
)

replace github.com/skinnyguy/spectra-services/core => ../core
