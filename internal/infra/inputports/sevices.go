package inputports

import (
	"go-clean-architecture-example/internal/app"
	"go-clean-architecture-example/internal/infra/inputports/http"
)

//Services contains the ports services
type Services struct {
	Server *http.Server
}

//NewServices instantiates the services of input ports
func NewServices(appServices app.Services) Services {
	return Services{Server: http.NewServer(appServices)}
}
