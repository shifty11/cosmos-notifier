package notifier

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/microcosm-cc/bluemonday"
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/log"
	"github.com/shifty11/cosmos-notifier/types"
	"regexp"
	"strconv"
)

type DiscordNotifier struct {
	discordBotToken       string
	discordChannelManager database.IDiscordChannelManager
	maxMsgLength          int
}

func newDiscordNotifier(managers *database.DbManagers, discordBotToken string) *DiscordNotifier {
	return &DiscordNotifier{
		discordBotToken:       discordBotToken,
		discordChannelManager: managers.DiscordChannelManager,
		maxMsgLength:          2000,
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

func (n *DiscordNotifier) notify(
	subscribedIds []types.DiscordChannelQueryResult,
	contractOrChainName string,
	proposalId int,
	proposalTitle string,
	proposalDescription string,
) {
	p := bluemonday.StripTagsPolicy()

	var textMsgs []string
	text := fmt.Sprintf("🎉  **%v - Proposal %v\n\n%v**\n\n*%v*",
		contractOrChainName,
		proposalId,
		p.Sanitize(proposalTitle),
		n.sanitizeUrls(p.Sanitize(proposalDescription)))
	if len(text) <= n.maxMsgLength {
		textMsgs = append(textMsgs, text)
	} else {
		textMsgs = append(textMsgs, text[:n.maxMsgLength-1]+"*")
		text = text[:len(text)-1] // remove the last character which is a *
		for _, chunk := range chunks(text[n.maxMsgLength-1:], n.maxMsgLength-2) {
			textMsgs = append(textMsgs, fmt.Sprintf("*%v*", chunk))
		}
	}

	session := n.startDiscordSession()
	defer n.closeDiscordSession(session)

	var errIds []int64
	for _, dc := range subscribedIds {
		log.Sugar.Debugf("Notifying discord channel %v (%v)", dc.Name, dc.ChannelId)
		for _, textMsg := range textMsgs {
			var _, err = session.ChannelMessageSendComplex(strconv.FormatInt(dc.ChannelId, 10), &discordgo.MessageSend{
				Content: textMsg,
			})
			if err != nil {
				if n.shouldDeleteUser(err) {
					errIds = append(errIds, dc.ChannelId)
				} else {
					log.Sugar.Errorf("Error sending proposal %v (%v) to discord channel %v (%v): %v", proposalId, contractOrChainName, dc.Name, dc.ChannelId, err)
				}
				break
			}
		}
	}

	if len(errIds) > 0 {
		n.discordChannelManager.DeleteMultiple(errIds)
	}
}

func (n *DiscordNotifier) broadcastMessage(ids []types.DiscordChannelQueryResult, message string) int {
	var textMsgs []string
	if len(message) <= n.maxMsgLength {
		textMsgs = append(textMsgs, message)
	} else {
		for _, chunk := range chunks(message, n.maxMsgLength) {
			textMsgs = append(textMsgs, chunk)
		}
	}

	session := n.startDiscordSession()
	defer n.closeDiscordSession(session)

	var errIds []int64
	for _, dc := range ids {
		log.Sugar.Debugf("Broadcasting message to discord channel %v (%v)", dc.Name, dc.ChannelId)
		for _, textMsg := range textMsgs {
			var _, err = session.ChannelMessageSendComplex(strconv.FormatInt(dc.ChannelId, 10), &discordgo.MessageSend{
				Content: textMsg,
			})
			if err != nil {
				if n.shouldDeleteUser(err) {
					errIds = append(errIds, dc.ChannelId)
				} else {
					log.Sugar.Errorf("Error while broadcasting message to discord channel %v (%v): %v", dc.Name, dc.ChannelId, err)
				}
				break
			}
		}
	}

	if len(errIds) > 0 {
		n.discordChannelManager.DeleteMultiple(errIds)
	}
	return len(errIds)
}

func (n *DiscordNotifier) sendVoteReminder(
	dc *ent.DiscordChannel,
	chainName string,
	proposalId int,
	proposalTitle string,
	remainingTimeText string,
) {
	log.Sugar.Debugf("Sending vote reminder for proposal %v on chain %v to discord channel %v (%v)", proposalId, chainName, dc.Name, dc.ChannelID)

	session := n.startDiscordSession()
	defer n.closeDiscordSession(session)

	p := bluemonday.StripTagsPolicy()

	var msgText string
	if remainingTimeText == "" {
		msgText = fmt.Sprintf("🗳️  **%v - Proposal Reminder %v\n\n%v**\n\nYou missed the voting deadline!",
			chainName,
			proposalId,
			p.Sanitize(proposalTitle),
		)
	} else {
		msgText = fmt.Sprintf("🗳️  **%v - Proposal Reminder %v\n\n%v**\n\nYou did not vote yet! You have **%v** left to vote.",
			chainName,
			proposalId,
			p.Sanitize(proposalTitle),
			remainingTimeText,
		)
	}

	var _, err = session.ChannelMessageSendComplex(
		strconv.FormatInt(dc.ChannelID, 10),
		&discordgo.MessageSend{Content: msgText})
	if err != nil {
		if n.shouldDeleteUser(err) {
			n.discordChannelManager.DeleteMultiple([]int64{dc.ChannelID})
		} else {
			log.Sugar.Errorf("Error sending vote reminder to discord channel %v (%v): %v", dc.Name, dc.ChannelID, err)
		}
	}
}
