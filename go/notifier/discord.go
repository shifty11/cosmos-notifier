package notifier

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/microcosm-cc/bluemonday"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/log"
	"regexp"
	"strconv"
)

type DiscordNotifier struct {
	discordBotToken       string
	discordChannelManager database.IDiscordChannelManager
}

func NewDiscordNotifier(managers *database.DbManagers, discordBotToken string) *DiscordNotifier {
	return &DiscordNotifier{
		discordBotToken:       discordBotToken,
		discordChannelManager: managers.DiscordChannelManager,
	}
}

func (n *DiscordNotifier) startDiscordSession() *discordgo.Session {
	var err error
	session, err := discordgo.New("Bot " + n.discordBotToken)
	if err != nil {
		log.Sugar.Fatalf("Invalid bot parameters: %v", err)
	}

	err = session.Open()
	if err != nil {
		log.Sugar.Fatalf("Cannot open the s: %v", err)
	}
	return session
}

func (n *DiscordNotifier) closeDiscordSession(session *discordgo.Session) {
	err := session.Close()
	if err != nil {
		log.Sugar.Errorf("Error while closing discord s: %v", err)
	}
}

func (n *DiscordNotifier) shouldDeleteUser(err error) bool {
	if restErr, ok := err.(*discordgo.RESTError); ok {
		return restErr.Response.StatusCode == 403 || restErr.Response.StatusCode == 404
	} else {
		return false
	}
}

func (n *DiscordNotifier) sanitizeUrls(text string) string {
	// Use <> around urls so that no embeds are created
	r, _ := regexp.Compile("https?://(www\\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9@:%_+.~#?&/=]*)")
	return r.ReplaceAllStringFunc(text,
		func(part string) string {
			return "<" + part + ">"
		},
	)
}

func (n *DiscordNotifier) Notify(entContract *ent.Contract, entProp *ent.Proposal) {
	p := bluemonday.StripTagsPolicy()

	var textMsgs []string
	text := fmt.Sprintf("🎉  **%v - Proposal %v\n\n%v**\n\n*%v*",
		entContract.Name,
		entProp.ProposalID,
		p.Sanitize(entProp.Title),
		n.sanitizeUrls(p.Sanitize(entProp.Description)),
	)
	if len(text) <= 2000 {
		textMsgs = append(textMsgs, text)
	} else {
		textMsgs = append(textMsgs, text[:1999]+"*")
		text = text[:len(text)-1] // remove the last character which is a *
		for _, chunk := range chunks(text[1999:], 1998) {
			textMsgs = append(textMsgs, fmt.Sprintf("*%v*", chunk))
		}
	}

	session := n.startDiscordSession()
	defer n.closeDiscordSession(session)

	var errIds []int64
	for _, dc := range n.discordChannelManager.GetSubscribedIds(entContract) {
		log.Sugar.Debugf("Notifying discord channel %v (%v)", dc.Name, dc.ChannelId)
		for _, textMsg := range textMsgs {
			var _, err = session.ChannelMessageSendComplex(strconv.FormatInt(dc.ChannelId, 10), &discordgo.MessageSend{
				Content: textMsg,
			})
			if err != nil {
				if n.shouldDeleteUser(err) {
					errIds = append(errIds, dc.ChannelId)
				} else {
					log.Sugar.Errorf("Error while sending proposal to discord channel %v (%v): %v", dc.Name, dc.ChannelId, err)
				}
				break
			}
		}
	}

	if len(errIds) > 0 {
		n.discordChannelManager.DeleteMultiple(errIds)
	}
}
