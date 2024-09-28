package handlers

import (
	"crazyfarmbot/config/di"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func SetupMainHandlerGroup(dispatcher *ext.Dispatcher, init *di.Initialization) {
	dispatcher.AddHandler(handlers.NewCommand("start", init.MainController.Start))
}
