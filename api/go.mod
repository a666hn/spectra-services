module github.com/skinnyguy/spectra-services/api

go 1.14

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/agnivade/levenshtein v1.1.0 // indirect
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/gorilla/websocket v1.4.2
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/rs/cors v1.6.0
	github.com/skinnyguy/spectra-services/core v0.0.0-20201107181655-47b666ccce71
	github.com/vektah/gqlparser/v2 v2.1.0
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace github.com/skinnyguy/spectra-services/core => ../core
