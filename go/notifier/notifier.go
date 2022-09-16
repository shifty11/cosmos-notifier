package notifier

import (
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/log"
)

type Notifier struct {
	telegramNotifier *TelegramNotifier
	discordNotifier  *DiscordNotifier
}

func NewNotifier(managers *database.DbManagers, telegramBotToken string, telegramEndpoint string, discordBotToken string) *Notifier {
	return &Notifier{
		telegramNotifier: NewTelegramNotifier(managers, telegramBotToken, telegramEndpoint),
		discordNotifier:  NewDiscordNotifier(managers, discordBotToken),
	}
}

func (n *Notifier) Notify(entContract *ent.Contract, entProp *ent.Proposal) {
	log.Sugar.Infof("Notifying for proposal %v on %v", entProp.ProposalID, entContract.Name)
	n.telegramNotifier.Notify(entContract, entProp)
	n.discordNotifier.Notify(entContract, entProp)
}

func chunks(text string, limit int) []string {
	if len(text) <= limit {
		return []string{text}
	}
	var result []string
	for len(text) > 0 {
		if len(text) < limit {
			limit = len(text)
		}
		result = append(result, text[:limit])
		text = text[limit:]
	}
	return result
}
