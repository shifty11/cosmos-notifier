package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/log"
	"os"
	"os/signal"
)

//goland:noinspection GoNameStartsWithPackageName
type DiscordClient struct {
	s                     *discordgo.Session
	discordChannelManager database.IDiscordChannelManager
	proposalManager       *database.ProposalManager
	botToken              string
	webAppUrl             string
}

func NewDiscordClient(managers *database.DbManagers, botToken string, webAppUrl string) *DiscordClient {
	return &DiscordClient{
		discordChannelManager: managers.DiscordChannelManager,
		proposalManager:       managers.ProposalManager,
		botToken:              botToken,
		webAppUrl:             webAppUrl,
	}
}

func (dc DiscordClient) initDiscord() *discordgo.Session {
	log.Sugar.Info("Init discord bot")

	var err error
	s, err := discordgo.New("Bot " + dc.botToken)
	if err != nil {
		log.Sugar.Fatalf("Invalid bot parameters: %v", err)
	}
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		go func() {
			switch i.Type {
			case discordgo.InteractionApplicationCommand:
				if h, ok := cmdHandlers[i.ApplicationCommandData().Name]; ok {
					h(dc, s, i)
				}
			}
		}()
	})
	return s
}

func (dc DiscordClient) addCommands() {
	for _, v := range cmds {
		_, err := dc.s.ApplicationCommandCreate(dc.s.State.User.ID, "", v)
		if err != nil {
			log.Sugar.Panic("Cannot create '%v' command: %v", v.Name, err)
		}
	}
}

func (dc DiscordClient) removeCommands() {
	registeredCommands, err := dc.s.ApplicationCommands(dc.s.State.User.ID, "")
	if err != nil {
		log.Sugar.Fatalf("Could not fetch registered commands: %v", err)
	}

	for _, v := range registeredCommands {
		err := dc.s.ApplicationCommandDelete(dc.s.State.User.ID, "", v.ID)
		if err != nil {
			log.Sugar.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}

func (dc DiscordClient) Start() {
	dc.s = dc.initDiscord()
	log.Sugar.Info("Start discord bot")

	err := dc.s.Open()
	if err != nil {
		log.Sugar.Fatalf("Cannot open the s: %v", err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer dc.s.Close()

	dc.removeCommands()
	dc.addCommands()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
