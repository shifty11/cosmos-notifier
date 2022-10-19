package notifier

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/microcosm-cc/bluemonday"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/types"
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

type TelegramNotifier struct {
	telegramApi         *tgbotapi.BotAPI
	telegramChatManager database.ITelegramChatManager
}

func NewTelegramNotifier(managers *database.DbManagers, telegramBotToken string, telegramEndpoint string) *TelegramNotifier {
	if telegramEndpoint == "" {
		telegramEndpoint = tgbotapi.APIEndpoint
	}
	telegramApi, err := tgbotapi.NewBotAPIWithAPIEndpoint(telegramBotToken, telegramEndpoint)
	if err != nil {
		log.Sugar.Panicf("Cannot create telegram bot: %v", err)
	}
	return &TelegramNotifier{
		telegramApi:         telegramApi,
		telegramChatManager: managers.TelegramChatManager,
	}
}

func (n *TelegramNotifier) shouldDeleteUser(err error) bool {
	if err != nil {
		return slices.Contains(forbiddenErrors, err.Error())
	}
	return false
}

func (n *TelegramNotifier) Notify(
	subscribedIds []types.TgChatQueryResult,
	contractOrChainName string,
	proposalId int,
	proposalTitle string,
	proposalDescription string,
) {
	p := bluemonday.StripTagsPolicy()

	var textMsgs []string
	text := fmt.Sprintf("🎉  <b>%v - Proposal %v\n\n%v</b>\n\n<i>%v</i>",
		contractOrChainName,
		proposalId,
		p.Sanitize(proposalTitle),
		p.Sanitize(proposalDescription),
	)
	if len(text) <= 4096 {
		textMsgs = append(textMsgs, text)
	} else {
		textMsgs = append(textMsgs, text[:4092]+"</i>")
		text = text[:len(text)-1] // remove the last character which is a *
		for _, chunk := range chunks(text[4092:], 4089) {
			textMsgs = append(textMsgs, fmt.Sprintf("<i>%v</i>", chunk))
		}
	}

	var errIds []int64
	for _, tg := range subscribedIds {
		log.Sugar.Debugf("Notifying telegram chat %v (%v)", tg.Name, tg.ChatId)
		msg := tgbotapi.NewMessage(tg.ChatId, text)
		msg.ParseMode = "html"
		msg.DisableWebPagePreview = true

		_, err := n.telegramApi.Send(msg)
		if err != nil {
			if n.shouldDeleteUser(err) {
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
