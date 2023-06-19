package main

import (
	"go-clean-architecture-example/internal/app"
	"go-clean-architecture-example/internal/infra/inputports"
	"go-clean-architecture-example/internal/infra/interfaceadapters"
	"go-clean-architecture-example/pkg/time"
	"go-clean-architecture-example/pkg/uuid"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	interfaceAdapterServices := interfaceadapters.NewServices()
	tp := time.NewTimeProvider()
	up := uuid.NewUUIDProvider()
	appServices := app.NewServices(interfaceAdapterServices.CragRepository, interfaceAdapterServices.NotificationService, up, tp)
	inputPortsServices := inputports.NewServices(appServices)
	inputPortsServices.Server.ListenAndServe(":8080")
}
