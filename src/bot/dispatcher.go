package bot

import (
	"crazyfarmbot/config/di"
	"crazyfarmbot/src/bot/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
)

func InitDispatcher(init *di.Initialization) *ext.Dispatcher {
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	{
		handlers.SetupMainHandlerGroup(dispatcher, init)
	}
	return dispatcher
}
