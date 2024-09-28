// go:build wireinject
//go:build wireinject
// +build wireinject

package di

import (
	"crazyfarmbot/config"
	"crazyfarmbot/src/controller"
	"crazyfarmbot/src/services"
	"github.com/google/wire"
)

var connectionsSet = wire.NewSet(
	config.GetBotInstance,
	config.ConnectToNatsBroker,
)

var natsBrokerSet = wire.NewSet(
	config.NatsBrokerInit,
	wire.Bind(new(config.NatsBroker), new(*config.NatsBrokerImpl)),
)

var serviceSet = wire.NewSet(
	services.SubServiceInit,
	wire.Bind(new(services.SubService), new(*services.SubServiceImpl)),
)
var controllersSet = wire.NewSet(
	controller.MainControllerInit,
	wire.Bind(new(controller.MainController), new(*controller.MainControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization,
		connectionsSet,
		controllersSet,
		natsBrokerSet,
		serviceSet)
	return nil
}
