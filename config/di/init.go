package di

import (
	"crazyfarmbot/config"
	"crazyfarmbot/src/controller"
	"crazyfarmbot/src/services"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

type Initialization struct {
	Bot            *gotgbot.Bot
	SubService     services.SubService
	MainController controller.MainController
	Nats           config.NatsBroker
}

func NewInitialization(
	Bot *gotgbot.Bot,
	SubService services.SubService,
	MainController controller.MainController,
	Nats config.NatsBroker) *Initialization {
	return &Initialization{
		Bot:            Bot,
		SubService:     SubService,
		MainController: MainController,
		Nats:           Nats,
	}
}
