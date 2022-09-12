package notifier

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/microcosm-cc/bluemonday"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/log"
	"golang.org/x/exp/slices"
)

var forbiddenErrors = []string{
	"Forbidden: bot was blocked by the user",
	"Forbidden: bot was kicked from the group chat",
	"Forbidden: bot was kicked from the supergroup chat",
	"Forbidden: bot is not a member of the supergroup chat",
	"Forbidden: user is deactivated",
	"Bad Request: chat not found",
}

func shouldDeleteUser(err error) bool {
	if err != nil {
		return slices.Contains(forbiddenErrors, err.Error())
	}
	return false
}

func (n *Notifier) notifyTelegram(entContract *ent.Contract, entProp *ent.Proposal) {
	p := bluemonday.StripTagsPolicy()

	text := fmt.Sprintf("🎉  <b>%v - Proposal %v\n\n%v</b>\n\n<i>%v</i>",
		entContract.Name,
		entProp.ProposalID,
		p.Sanitize(entProp.Title),
		p.Sanitize(entProp.Description),
	)
	if len(text) > 4096 {
		text = text[:4088] + "</i> ..."
	}

	var errIds []int64
	for _, tg := range n.telegramChatManager.GetSubscribedIds(entContract) {
		log.Sugar.Debugf("Notifying telegram chat %v (%v)", tg.Name, tg.ChatId)
		msg := tgbotapi.NewMessage(tg.ChatId, text)
		msg.ParseMode = "html"
		msg.DisableWebPagePreview = true

		_, err := n.telegramApi.Send(msg)
		if err != nil {
			if shouldDeleteUser(err) {
				errIds = append(errIds, tg.ChatId)
			} else {
				log.Sugar.Errorf("Error sending telegram message to %v (%v): %v", tg.Name, tg.ChatId, err)
			}
		}
	}

	if len(errIds) > 0 {
		n.telegramChatManager.DeleteMultiple(errIds)
	}
}
