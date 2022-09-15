package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/log"
)

func (client TelegramClient) handleCommand(update *tgbotapi.Update) {
	switch MessageCommand(update.Message.Command()) { // Check for non admin commands
	case MessageCmdStart, MessageCmdSubscriptions:
		client.handleStart(update)
	}
}

// groups -> just admins and creators can interact with the bot
// private -> everything is allowed
func (client TelegramClient) isInteractionAllowed(update *tgbotapi.Update) bool {
	if isGroupX(update) {
		return client.isUpdateFromCreatorOrAdministrator(update)
	}
	return true
}

// Handles updates for only 1 user in a serial way
func (client TelegramClient) handleUpdates(channel chan tgbotapi.Update) {
	for update := range channel {
		chatId := getChatIdX(&update)
		if client.isInteractionAllowed(&update) {
			if update.Message != nil && update.Message.IsCommand() {
				client.handleCommand(&update)
			} else if update.CallbackQuery != nil {
				log.Sugar.Errorf("Callback query not implemented: %v", update.CallbackQuery)
			}
		} else {
			log.Sugar.Debugf("Interaction with bot for user #%v is not allowed", chatId)
			if update.CallbackQuery != nil {
				log.Sugar.Errorf("Callback query not implemented: %v", update.CallbackQuery)
			}
		}
		client.updateCountChannel <- UpdateCount{ChatId: chatId, Updates: -1}
	}
}

type UpdateCount struct {
	ChatId  int64
	Updates int
}

func (client TelegramClient) hasChannel(channelId int64) bool {
	for key := range client.updateChannels {
		if key == channelId {
			return true
		}
	}
	return false
}

func (client TelegramClient) sendToChannelAsync(chatId int64, update tgbotapi.Update) {
	client.updateCountChannel <- UpdateCount{ChatId: chatId, Updates: 1}
	client.updateChannels[chatId] <- update
}

func (client TelegramClient) sendToChannel(update *tgbotapi.Update) {
	chatId := getChatIdX(update)
	if !client.hasChannel(chatId) {
		client.updateChannels[chatId] = make(chan tgbotapi.Update)
		go client.handleUpdates(client.updateChannels[chatId])
	}
	go client.sendToChannelAsync(chatId, *update)
}

// Keeps track of all the user channels and closes them if there are no more updates
func (client TelegramClient) manageUpdateChannels() {
	var count = make(map[int64]int)
	for msg := range client.updateCountChannel {
		count[msg.ChatId] += msg.Updates
		if count[msg.ChatId] == 0 {
			close(client.updateChannels[msg.ChatId])
			delete(client.updateChannels, msg.ChatId)
			delete(count, msg.ChatId)
		}
	}
}

//goland:noinspection GoNameStartsWithPackageName
type TelegramClient struct {
	api *tgbotapi.BotAPI

	// updateChannels contains one update channel for every user.
	// This means the updates can be processed parallel for multiple users but serial for every single user
	updateChannels map[int64]chan tgbotapi.Update

	// updateCountChannel is used to communicate to `manageUpdateChannels` from `handleUpdates`
	updateCountChannel chan UpdateCount

	UserManager         *database.UserManager
	TelegramChatManager database.ITelegramChatManager
	ProposalManager     *database.ProposalManager

	telegramToken    string
	telegramEndpoint string
	webAppUrl        string
}

func NewTelegramClient(
	managers *database.DbManagers,
	telegramToken string,
	telegramEndpoint string,
	webAppUrl string,
) *TelegramClient {
	if telegramEndpoint == "" {
		telegramEndpoint = tgbotapi.APIEndpoint
	}
	api, err := tgbotapi.NewBotAPIWithAPIEndpoint(telegramToken, telegramEndpoint)
	if err != nil {
		log.Sugar.Panic(err)
	}
	return &TelegramClient{
		api:                api,
		updateChannels:     make(map[int64]chan tgbotapi.Update),
		updateCountChannel: make(chan UpdateCount),

		UserManager:         managers.UserManager,
		TelegramChatManager: managers.TelegramChatManager,
		ProposalManager:     managers.ProposalManager,

		telegramToken:    telegramToken,
		telegramEndpoint: telegramEndpoint,
		webAppUrl:        webAppUrl,
	}
}

func (client TelegramClient) Start() {
	log.Sugar.Info("Start telegram bot")

	updateConfig := tgbotapi.NewUpdate(0)
	updates := client.api.GetUpdatesChan(updateConfig)

	go client.manageUpdateChannels()

	for update := range updates {
		if !hasChatId(&update) { // no chat id means there is something strange or the update is not for us
			continue
		}

		client.sendToChannel(&update)
	}
}
