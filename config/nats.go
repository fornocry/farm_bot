package config

import (
	"crazyfarmbot/src/services"
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type NatsBroker interface {
	CheckSubscribe()
}

type NatsBrokerImpl struct {
	nc         *nats.Conn
	subService services.SubService
}

func (u NatsBrokerImpl) CheckSubscribe() {
	_, err := u.nc.Subscribe("check_subscribe", func(m *nats.Msg) {
		parts := strings.Split(string(m.Data), ",")
		if len(parts) != 2 {
			log.Error("Invalid request format. Expected 2 parts.")
			return
		}
		userId, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Error("Failed to parse User ID. Error: ", err)
			return
		}
		channelId, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Error("Failed to parse Channel ID. Error: ", err)
			return
		}
		isSubs := u.subService.CheckSubscribed(channelId, userId)
		if isSubs {
			m.Respond([]byte("1"))
		} else {
			m.Respond([]byte("0"))
		}
	})
	if err != nil {
		log.Error("Failed to subscribe. Error: ", err)
		return
	}
}

func (u NatsBrokerImpl) Init() {
	log.Infoln("Initialization nats...")
	u.CheckSubscribe()
	log.Infoln("Initialized nats success")
}

func NatsBrokerInit(nc *nats.Conn,
	subService services.SubService) *NatsBrokerImpl {
	broker := &NatsBrokerImpl{
		nc:         nc,
		subService: subService,
	}
	broker.Init()
	return broker
}
