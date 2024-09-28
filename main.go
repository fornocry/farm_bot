package main

import (
	"crazyfarmbot/config"
	"crazyfarmbot/config/di"
	"crazyfarmbot/src/bot"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

func main() {
	init := di.Init()
	botInst := init.Bot
	dispatcher := bot.InitDispatcher(init)
	err := bot.StartUpdater(dispatcher, botInst)
	if err != nil {
		log.Fatal(err)
	}
}
