package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/icza/gog"
	"github.com/shifty11/dao-dao-notifier/log"
)

type MessageCommand string

const (
	MessageCmdStart         MessageCommand = "start"
	MessageCmdSubscriptions MessageCommand = "subscriptions"
	MessageCmdStop          MessageCommand = "stop"
)

func (client TelegramClient) handleCommand(update *tgbotapi.Update) {
	switch MessageCommand(update.Message.Command()) {
	case MessageCmdStart, MessageCmdSubscriptions:
		client.handleStart(update)
	case MessageCmdStop:
		client.handleStop(update)
	}
}

const subscriptionsMsg = `🚀 DaoDao Notifier started.
%v
🔔 Active subscriptions: %v

*How does it work?*
- You subscribe to a DAO (ex: [rawdao](https://www.rawdao.zone/vote))
- A member of this DAO creates a governance proposal
- A notification that a new proposal is up for voting is sent to this chat

To stop the bot send the command /stop
`

func (client TelegramClient) handleStart(update *tgbotapi.Update) {
	userId := getUserIdX(update)
	userName := getUserName(update)
	chatId := getChatIdX(update)
	chatName := getChatName(update)
	isGroup := isGroupX(update)

	client.TelegramChatManager.CreateOrUpdateChat(userId, userName, chatId, chatName, isGroup)
	client.TelegramChatManager.MigrateOldUsers(userId)

	adminText := ""
	if isGroup {
		adminText += "\n👮‍♂ Bot admins in this chat\n"
		for _, user := range client.TelegramChatManager.GetChatUsers(chatId) {
			adminText += fmt.Sprintf("- @%v\n", user.Name)
		}
	}

	log.Sugar.Debugf("Send start to %v %v (%v)", gog.If(isGroup, "group", "user"), chatName, chatId)

	var buttons [][]Button
	buttons = append(buttons, client.getSubscriptionButtonRow(update))
	replyMarkup := createKeyboard(buttons)

	cnt := client.TelegramChatManager.CountSubscriptions(chatId)
	msg := tgbotapi.NewMessage(chatId, fmt.Sprintf(subscriptionsMsg, adminText, cnt))
	msg.ReplyMarkup = replyMarkup
	msg.ParseMode = "markdown"
	msg.DisableWebPagePreview = true
	_, err := client.api.Send(msg)
	if err != nil {
		log.Sugar.Errorf("Error while sending /start response for user %v (%v): %v", chatName, chatId, err)
	}
}

func (client TelegramClient) getSubscriptionButtonRow(update *tgbotapi.Update) []Button {
	var buttonRow []Button
	button := NewButton("🔔 Subscriptions")
	button.LoginURL = &tgbotapi.LoginURL{URL: fmt.Sprintf("%v?chat-id=%v", client.webAppUrl, getChatIdX(update)), RequestWriteAccess: true}
	buttonRow = append(buttonRow, button)
	return buttonRow
}

func (client TelegramClient) handleStop(update *tgbotapi.Update) {
	userId := getUserIdX(update)
	chatId := getChatIdX(update)
	chatName := getChatName(update)
	isGroup := isGroupX(update)

	//goland:noinspection GoUnhandledErrorResult
	client.TelegramChatManager.Delete(userId, chatId)

	log.Sugar.Debugf("Send stop to %v %v (%v)", gog.If(isGroup, "group", "user"), chatName, chatId)

	var buttons [][]Button
	buttons = append(buttons, client.getSubscriptionButtonRow(update))
	replyMarkup := createKeyboard(buttons)

	msg := tgbotapi.NewMessage(chatId, "😴 Bot stopped. Send /start to start it again.")
	msg.ReplyMarkup = replyMarkup
	msg.ParseMode = "markdown"
	msg.DisableWebPagePreview = true
	_, err := client.api.Send(msg)
	if err != nil {
		log.Sugar.Errorf("Error while sending /stop response for user %v (%v): %v", chatName, chatId, err)
	}
}
