package discord

import (
	"encoding/base64"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/icza/gog"
	"github.com/shifty11/dao-dao-notifier/log"
	"net/url"
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

			state := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("chat-id=%v", channelId)))
			params := url.Values{}
			params.Add("client_id", dc.clientId)
			params.Add("redirect_uri", dc.webAppUrl)
			params.Add("response_type", "code")
			params.Add("scope", "identify")
			params.Add("state", state)
			redirectUrl := fmt.Sprintf("https://discord.com/oauth2/authorize?%v", params.Encode())
			dc.discordChannelManager.CreateOrUpdateChannel(userId, userName, channelId, channelName, isGroup)
			dc.discordChannelManager.MigrateOldUsers(userId)
			cntSubs := dc.discordChannelManager.CountSubscriptions(channelId)

			text := ""
			if isGroup {
				adminText := ""
				for _, user := range dc.discordChannelManager.GetChannelUsers(channelId) {
					adminText += fmt.Sprintf("- `%v`\n", user.Name)
				}
				text = fmt.Sprintf(":rocket: DaoDao Notifier started.\n\n") +
					fmt.Sprintf(":police_officer: Bot admins in this channel:\n%v\n", adminText) +
					fmt.Sprintf(":bell: Active subscriptions: %v\n\n", cntSubs) +
					fmt.Sprintf("Go to **[DaoDao Notifier](%v)** to change subscriptions for this channel.\n\n", redirectUrl) +
					"**How does it work?**\n" +
					"- You subscribe this channel to a DAO (ex: [rawdao](https://www.rawdao.zone/vote))\n" +
					"- A member of this DAO creates a governance proposal\n" +
					"- A notification that a new proposal is up for voting is sent to this channel\n\n" +
					"To register another user as admin he has to send the command `/start` to the bot.\n" +
					"To stop the bot send the command `/stop`."
			} else {
				text = fmt.Sprintf(":rocket: DaoDao Notifier started.\n\n") +
					fmt.Sprintf(":bell: Active subscriptions: %v\n\n", cntSubs) +
					fmt.Sprintf("Go to **[DaoDao Notifier](%v)** to change your subscriptions.\n\n", redirectUrl) +
					"**How does it work?**\n" +
					"- You subscribe to a DAO (ex: [rawdao](https://www.rawdao.zone/vote))\n" +
					"- A member of this DAO creates a governance proposal\n" +
					"- A notification that a new proposal is up for voting is sent to you\n\n" +
					"To stop the bot send the command `/stop`."
			}

			log.Sugar.Debugf("Send start to %v %v (%v)", gog.If(isGroup, "group", "user"), channelName, channelId)

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
			channelName := getChannelName(s, i)
			isGroup := isGroup(i)

			//goland:noinspection GoUnhandledErrorResult
			dc.discordChannelManager.Delete(userId, channelId)
			text := ":sleeping: Bot stopped. Send `/start` to start it again."

			log.Sugar.Debugf("Send stop to %v %v (%v)", gog.If(isGroup, "group", "user"), channelName, channelId)

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
