package controller

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type MainController interface {
	Start(b *gotgbot.Bot, ctx *ext.Context) error
}

type MainControllerImpl struct {
}

func (u MainControllerImpl) Start(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("<b>Welcome to Farm⭐️ \n\n Start growing asap! Press Farm!</b>"), &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}

func MainControllerInit() *MainControllerImpl {
	return &MainControllerImpl{}
}
