package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shifty11/dao-dao-notifier/log"
)

const subscriptionsMsg = `🔔 *Subscriptions*

Select the DAO's that you want to follow. You will receive notifications about new governance proposals.
`

func (client TelegramClient) handleSubscription(update *tgbotapi.Update) {
	userId := getUserIdX(update)
	userName := getUserName(update)
	chatId := getChatIdX(update)
	chatName := getChatName(update)
	isGroup := isGroupX(update)

	client.TelegramChatManager.UpdateOrCreateChat(userId, userName, chatId, chatName, isGroup)

	if update.Message != nil && update.Message.Chat != nil && update.Message.Chat.Type == "group" {
		log.Sugar.Debugf("Send subscriptions to group '%v' #%v", update.Message.Chat.Title, chatId)
	} else {
		log.Sugar.Debugf("Send subscriptions to user #%v", chatId)
	}

	var buttons [][]Button
	buttons = append(buttons, client.getSubscriptionButtonRow(update))
	replyMarkup := createKeyboard(buttons)

	msg := tgbotapi.NewMessage(chatId, subscriptionsMsg)
	msg.ReplyMarkup = replyMarkup
	msg.ParseMode = "markdown"
	msg.DisableWebPagePreview = true
	_, err := client.api.Send(msg)
	if err != nil {
		log.Sugar.Errorf("Error while sending login button for user #%v: %v", chatId, err)
	}
}

func (client TelegramClient) getSubscriptionButtonRow(update *tgbotapi.Update) []Button {
	var buttonRow []Button
	button := NewButton("🌐 Subscriptions")
	button.LoginURL = &tgbotapi.LoginURL{URL: fmt.Sprintf("%v?chat_id=%v", client.webAppUrl, getChatIdX(update)), RequestWriteAccess: true}
	buttonRow = append(buttonRow, button)
	return buttonRow
}
