package bot

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func StartUpdater(dispatcher *ext.Dispatcher, bot *gotgbot.Bot) error {
	updater := ext.NewUpdater(dispatcher, nil)
	err := updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: 10 * time.Second,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to start polling: %w", err) // Return error instead of panic
	}
	log.Printf("@%s has been started...\n", bot.User.Username)
	updater.Idle()

	return nil
}
