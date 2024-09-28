package config

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func ConnectToDB() *gorm.DB {
	log.Infoln("Connecting to database...")
	var err error
	dbDsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(postgres.Open(dbDsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
	}

	return db
}

func ConnectToNatsBroker() *nats.Conn {
	log.Infoln("Connecting to nats...")
	natsDsn := os.Getenv("NATS_DSN")
	nc, _ := nats.Connect(natsDsn)
	return nc
}

func GetBotInstance() *gotgbot.Bot {
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		panic("TOKEN environment variable is empty")
	}

	bot, err := gotgbot.NewBot(token, nil)
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}
	return bot
}
