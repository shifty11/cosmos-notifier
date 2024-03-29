package notifier

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/microcosm-cc/bluemonday"
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/log"
	"github.com/shifty11/cosmos-notifier/types"
	"golang.org/x/exp/slices"
)

var forbiddenErrors = []string{
	"Forbidden: bot was blocked by the user",
	"Forbidden: bot was kicked from the group chat",
	"Forbidden: bot was kicked from the supergroup chat",
	"Forbidden: bot is not a member of the supergroup chat",
	"Forbidden: user is deactivated",
	"Bad Request: chat not found",
	"Bad Request: group chat was upgraded to a supergroup chat",
}

type telegramNotifier struct {
	telegramApi         *tgbotapi.BotAPI
	telegramChatManager database.ITelegramChatManager
	maxMsgLength        int
}

func newTelegramNotifier(managers *database.DbManagers, telegramBotToken string, telegramEndpoint string) *telegramNotifier {
	if telegramEndpoint == "" {
		telegramEndpoint = tgbotapi.APIEndpoint
	}
	telegramApi, err := tgbotapi.NewBotAPIWithAPIEndpoint(telegramBotToken, telegramEndpoint)
	if err != nil {
		log.Sugar.Panicf("Cannot create telegram bot: %v", err)
	}
	return &telegramNotifier{
		telegramApi:         telegramApi,
		telegramChatManager: managers.TelegramChatManager,
		maxMsgLength:        4096,
	}
}

func (n *telegramNotifier) shouldDeleteUser(err error) bool {
	if err != nil {
		return slices.Contains(forbiddenErrors, err.Error())
	}
	return false
}

func (n *telegramNotifier) notify(
	subscribedIds []types.TgChatQueryResult,
	contractOrChainName string,
	proposalId int,
	proposalTitle string,
	proposalDescription string,
) {
	p := bluemonday.StripTagsPolicy()

	var textMsgs []string
	message := fmt.Sprintf("🎉  <b>%v - Proposal %v\n\n%v</b>\n\n<i>%v</i>",
		contractOrChainName,
		proposalId,
		p.Sanitize(proposalTitle),
		p.Sanitize(proposalDescription),
	)
	if len(message) <= n.maxMsgLength {
		textMsgs = append(textMsgs, message)
	} else {
		textMsgs = append(textMsgs, message[:n.maxMsgLength-4]+"</i>")
		message = message[:len(message)-4] // remove the last 4 characters which are "</i>"
		for _, chunk := range chunks(message[n.maxMsgLength-4:], n.maxMsgLength-7) {
			textMsgs = append(textMsgs, fmt.Sprintf("<i>%v</i>", chunk))
		}
	}

	var errIds []int64
	for _, tg := range subscribedIds {
		log.Sugar.Debugf("Notifying telegram chat %v (%v)", tg.Name, tg.ChatId)
		for _, textMsg := range textMsgs {
			msg := tgbotapi.NewMessage(tg.ChatId, textMsg)
			msg.ParseMode = "html"
			msg.DisableWebPagePreview = true

			_, err := n.telegramApi.Send(msg)
			if err != nil {
				if n.shouldDeleteUser(err) {
					errIds = append(errIds, tg.ChatId)
				} else {
					log.Sugar.Errorf("Error sending proposal %v (%v) to telegram chat %v (%v): %v", proposalId, contractOrChainName, tg.Name, tg.ChatId, err)
				}
				break
			}
		}
	}

	if len(errIds) > 0 {
		n.telegramChatManager.DeleteMultiple(errIds)
	}
}

func (n *telegramNotifier) broadcastMessage(ids []types.TgChatQueryResult, message string) int {
	var textMsgs []string
	if len(message) <= n.maxMsgLength {
		textMsgs = append(textMsgs, message)
	} else {
		for _, chunk := range chunks(message, n.maxMsgLength) {
			textMsgs = append(textMsgs, chunk)
		}
	}

	var errIds []int64
	for _, tg := range ids {
		log.Sugar.Debugf("Broadcasting message to telegram chat %v (%v)", tg.Name, tg.ChatId)
		msg := tgbotapi.NewMessage(tg.ChatId, message)
		msg.ParseMode = "html"
		msg.DisableWebPagePreview = true

		_, err := n.telegramApi.Send(msg)
		if err != nil {
			if n.shouldDeleteUser(err) {
				errIds = append(errIds, tg.ChatId)
			} else {
				log.Sugar.Errorf("Error broadcasting message to telegram chat %v (%v): %v", tg.Name, tg.ChatId, err)
			}
		}
	}

	if len(errIds) > 0 {
		n.telegramChatManager.DeleteMultiple(errIds)
	}
	return len(errIds)
}

func (n *telegramNotifier) sendVoteReminder(
	telegramChat *ent.TelegramChat,
	chainName string,
	proposalId int,
	proposalTitle string,
	voter string,
	remainingTimeText string,
) {
	log.Sugar.Debugf("Sending vote reminder for proposal %v on chain %v to telegram chat %v (%v)", proposalId, chainName, telegramChat.Name, telegramChat.ChatID)

	var textMsg string
	if remainingTimeText == "" {
		textMsg = fmt.Sprintf("🗳️  <b>%v - Reminder Proposal %v\n\n%v</b>\n\n<b>%v</b> missed the voting deadline!",
			chainName,
			proposalId,
			proposalTitle,
			voter,
		)

	} else {
		textMsg = fmt.Sprintf("🗳️  <b>%v - Reminder Proposal %v\n\n%v</b>\n\n<b>%v</b> did not vote yet!\n<b>%v</b> until the vote ends.",
			chainName,
			proposalId,
			proposalTitle,
			voter,
			remainingTimeText,
		)
	}

	msg := tgbotapi.NewMessage(telegramChat.ChatID, textMsg)
	msg.ParseMode = "html"
	msg.DisableWebPagePreview = true

	_, err := n.telegramApi.Send(msg)
	if err != nil {
		if n.shouldDeleteUser(err) {
			n.telegramChatManager.DeleteMultiple([]int64{telegramChat.ChatID})
		} else {
			log.Sugar.Errorf("Error sending vote reminder to telegram chat %v (%v): %v", telegramChat.Name, telegramChat.ChatID, err)
		}
	}
}
