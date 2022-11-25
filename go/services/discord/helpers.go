package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/shifty11/cosmos-notifier/log"
	"regexp"
	"strconv"
)

func getChannelId(i *discordgo.InteractionCreate) int64 {
	channelId, err := strconv.ParseInt(i.ChannelID, 10, 64)
	if err != nil {
		log.Sugar.Panicf("Error while converting channel ID to int: %v", err)
	}
	return channelId
}

func getUserIdX(i *discordgo.InteractionCreate) int64 {
	if isGroup(i) {
		userId, err := strconv.ParseInt(i.Member.User.ID, 10, 64)
		if err != nil {
			log.Sugar.Panicf("Error while converting user ID to int: %v", err)
		}
		return userId
	} else {
		userId, err := strconv.ParseInt(i.User.ID, 10, 64)
		if err != nil {
			log.Sugar.Panicf("Error while converting user ID to int: %v", err)
		}
		return userId
	}
}

func getUserName(i *discordgo.InteractionCreate) string {
	if isGroup(i) {
		return i.Member.User.Username
	}
	return i.User.Username
}

func getChannelName(s *discordgo.Session, i *discordgo.InteractionCreate) string {
	if isGroup(i) {
		channel, err := s.Channel(i.ChannelID)
		if err != nil {
			log.Sugar.Errorf("Error while getting channel: %v", err)
		}
		return channel.Name
	}
	return i.User.Username
}

func isGroup(i *discordgo.InteractionCreate) bool {
	return i.Member != nil
}

func canInteractWithBot(s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	channel, err := s.Channel(i.ChannelID)
	if err != nil {
		log.Sugar.Debugf("Error while getting channel: %v", err)
		return false
	}
	if channel.Type == discordgo.ChannelTypeDM {
		return true
	}

	p, err := s.UserChannelPermissions(i.Interaction.Member.User.ID, i.ChannelID)
	if err != nil {
		log.Sugar.Errorf("Error while getting permissions: %v", err)
		return false
	}

	permAdmin := int64(discordgo.PermissionAdministrator)
	permManChan := int64(discordgo.PermissionManageChannels)
	permManServ := int64(discordgo.PermissionManageServer)
	return p&permAdmin == permAdmin || p&permManChan == permManChan || p&permManServ == permManServ
}

func sendEmptyResponse(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})
	if err != nil {
		log.Sugar.Errorf("Error while sending empty response: %v", err)
	}
}

func sanitizeUrls(text string) string {
	// Use <> around urls so that no embeds are created
	r, _ := regexp.Compile("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9@:%_\\+.~#?&//=]*)")
	return r.ReplaceAllStringFunc(text,
		func(part string) string {
			return "<" + part + ">"
		},
	)
}
