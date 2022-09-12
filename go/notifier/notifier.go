package notifier

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/log"
)

type Notifier struct {
	telegramApi *tgbotapi.BotAPI
}

func NewNotifier(telegramToken string, telegramEndpoint string) *Notifier {
	if telegramEndpoint == "" {
		telegramEndpoint = tgbotapi.APIEndpoint
	}
	telegramApi, err := tgbotapi.NewBotAPIWithAPIEndpoint(telegramToken, telegramEndpoint)
	if err != nil {
		log.Sugar.Panic(err)
	}
	return &Notifier{telegramApi: telegramApi}
}

func (n *Notifier) Notify(entProp *ent.Proposal) {
	log.Sugar.Infof("Notifying for proposal %v", entProp.ID)
	//TODO:	Implement
}
