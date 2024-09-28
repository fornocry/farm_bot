package services

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	log "github.com/sirupsen/logrus"
)

type SubService interface {
	CheckSubscribed(chatId int64, userId int64) bool
}

type SubServiceImpl struct {
	bot *gotgbot.Bot
}

func (u SubServiceImpl) CheckSubscribed(chatId int64, userId int64) bool {
	chat_member, err := u.bot.GetChatMember(chatId, userId, nil)
	if err != nil {
		log.Error(err)
		return false
	}
	return chat_member.GetStatus() != "left"
}

func SubServiceInit(bot *gotgbot.Bot) *SubServiceImpl {
	return &SubServiceImpl{
		bot: bot,
	}
}
