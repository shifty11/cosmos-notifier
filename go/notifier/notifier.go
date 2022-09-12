package notifier

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/log"
)

type Notifier struct {
	telegramApi         *tgbotapi.BotAPI
	telegramChatManager *database.TelegramChatManager
}

func NewNotifier(managers *database.DbManagers, telegramToken string, telegramEndpoint string) *Notifier {
	if telegramEndpoint == "" {
		telegramEndpoint = tgbotapi.APIEndpoint
	}
	telegramApi, err := tgbotapi.NewBotAPIWithAPIEndpoint(telegramToken, telegramEndpoint)
	if err != nil {
		log.Sugar.Panic(err)
	}
	return &Notifier{telegramApi: telegramApi, telegramChatManager: managers.TelegramChatManager}
}

func (n *Notifier) Notify(entContract *ent.Contract, entProp *ent.Proposal) {
	log.Sugar.Infof("Notifying for proposal %v on %v", entProp.ProposalID, entContract.Name)
	n.notifyTelegram(entContract, entProp)
}
