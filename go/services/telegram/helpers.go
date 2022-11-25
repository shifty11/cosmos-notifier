package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shifty11/cosmos-notifier/log"
	"golang.org/x/exp/slices"
)

type Button struct {
	Text     string
	LoginURL *tgbotapi.LoginURL
	WebApp   *tgbotapi.WebAppInfo
}

func NewButton(text string) Button {
	return Button{Text: text}
}

func createKeyboard(buttons [][]Button) *tgbotapi.InlineKeyboardMarkup {
	var keyboard [][]tgbotapi.InlineKeyboardButton
	for _, row := range buttons {
		var keyboardRow []tgbotapi.InlineKeyboardButton
		for _, button := range row {
			btn := tgbotapi.InlineKeyboardButton{Text: button.Text, LoginURL: button.LoginURL, WebApp: button.WebApp}
			keyboardRow = append(keyboardRow, btn)
		}
		keyboard = append(keyboard, keyboardRow)
	}
	return &tgbotapi.InlineKeyboardMarkup{InlineKeyboard: keyboard}
}

func hasChatId(update *tgbotapi.Update) bool {
	if update.CallbackQuery != nil {
		return true
	}
	if update.Message != nil {
		return true
	}
	return false
}

func getChatIdX(update *tgbotapi.Update) int64 {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.Message.Chat.ID
	}
	if update.Message != nil {
		return update.Message.Chat.ID
	}
	log.Sugar.Panic("getChatIdX: unreachable code reached!!!")
	return 0
}

func getChatName(update *tgbotapi.Update) string {
	if update.CallbackQuery != nil {
		if isGroupX(update) {
			return update.CallbackQuery.Message.Chat.Title
		}
		return update.CallbackQuery.Message.Chat.UserName
	}
	if update.Message != nil {
		if isGroupX(update) {
			return update.Message.Chat.Title
		}
		return update.Message.Chat.UserName
	}
	return ""
}

func isGroupX(update *tgbotapi.Update) bool {
	if update.CallbackQuery != nil {
		return !update.CallbackQuery.Message.Chat.IsPrivate()
	}
	if update.Message != nil {
		return !update.Message.Chat.IsPrivate()
	}
	log.Sugar.Panic("isGroupX: unreachable code reached!!!")
	return false
}

func getUserIdX(update *tgbotapi.Update) int64 {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.From.ID
	}
	if update.Message != nil {
		return update.Message.From.ID
	}
	log.Sugar.Panic("getUserIdX: unreachable code reached!!!")
	return 0
}

func getUserName(update *tgbotapi.Update) string {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.From.UserName
	}
	if update.Message != nil {
		return update.Message.From.UserName
	}
	return "<not found>"
}

var forbiddenErrors = []string{
	"Forbidden: bot was blocked by the user",
	"Forbidden: bot was kicked from the group chat",
	"Forbidden: bot was kicked from the supergroup chat",
	"Forbidden: bot is not a member of the supergroup chat",
	"Forbidden: user is deactivated",
	"Bad Request: chat not found",
}

func (client TelegramClient) isUpdateFromCreatorOrAdministrator(update *tgbotapi.Update) bool {
	chatId := getChatIdX(update)
	userId := getUserIdX(update)
	memberConfig := tgbotapi.GetChatMemberConfig{
		ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
			ChatID:             chatId,
			SuperGroupUsername: "",
			UserID:             userId,
		},
	}
	member, err := client.api.GetChatMember(memberConfig)
	if err != nil {
		if slices.Contains(forbiddenErrors, err.Error()) {
			log.Sugar.Debugf("Error while getting member (ChatID: %v; UserID: %v): %v", chatId, userId, err)
			return false
		}
		log.Sugar.Errorf("Error while getting member (ChatID: %v; UserID: %v): %v", chatId, userId, err)
		return false
	}
	return member.Status == "creator" || member.Status == "administrator"
}
