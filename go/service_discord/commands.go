package discord

import (
	"encoding/base64"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/shifty11/dao-dao-notifier/log"
)

var (
	startCmd = "start"
	stopCmd  = "stop"
	cmds     = []*discordgo.ApplicationCommand{
		{
			Name:        startCmd,
			Description: "Start the bot and receive notifications",
		},
		{
			Name:        stopCmd,
			Description: "Stop the bot",
		},
	}
	cmdHandlers = map[string]func(dc DiscordClient, s *discordgo.Session, i *discordgo.InteractionCreate){
		startCmd: func(dc DiscordClient, s *discordgo.Session, i *discordgo.InteractionCreate) {
			if !canInteractWithBot(s, i) {
				sendEmptyResponse(s, i)
				return
			}

			userId := getUserIdX(i)
			userName := getUserName(i)
			channelId := getChannelId(i)
			channelName := getChannelName(s, i)
			isGroup := isGroup(i)

			chatIdQueryParam := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("chat_id=%v", channelId)))
			// TODO: create redirect url from WEBAPP_URL
			url := "https://discord.com/oauth2/authorize?client_id=1018913065644867677&redirect_uri=http%3A%2F%2Flocalhost%3A40001&response_type=code&scope=identify&state=" + chatIdQueryParam
			// TODO: check if user has permissions in this channel
			_, created := dc.discordChannelManager.CreateOrUpdateChannel(userId, userName, channelId, channelName, isGroup)
			text := "Go to " +
				fmt.Sprintf("[DaoDao Notifier](%v)", url) +
				" to change your subscriptions.\n" +
				"You will then receive notifications about new proposals."
			if created {
				text = "Bot started.\n\n" + text
			} else {
				text = "Bot already started.\n\n" + text
			}

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: text,
				},
			})
			if err != nil {
				log.Sugar.Errorf("Error while sending subscriptions: %v", err)
			}
		},
		stopCmd: func(dc DiscordClient, s *discordgo.Session, i *discordgo.InteractionCreate) {
			if !canInteractWithBot(s, i) {
				sendEmptyResponse(s, i)
				return
			}

			userId := getUserIdX(i)
			channelId := getChannelId(i)

			dc.discordChannelManager.Delete(userId, channelId)
			text := "Bot stopped. Send /start to start it again."

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: text,
				},
			})
			if err != nil {
				log.Sugar.Errorf("Error while sending subscriptions: %v", err)
			}
		},
	}
)
