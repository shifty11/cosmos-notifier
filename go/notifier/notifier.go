package notifier

import (
	"errors"
	"fmt"
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/log"
	pb "github.com/shifty11/cosmos-notifier/services/grpc/protobuf/admin_service"
	"time"
)

type ContractNotifier interface {
	Notify(entContract *ent.Contract, entProp *ent.ContractProposal)
}

type ChainNotifier interface {
	Notify(entChain *ent.Chain, entProp *ent.ChainProposal)
	SendVoteReminder(data database.AddressTrackerWithChainProposal)
}

type GeneralNotifier interface {
	BroadcastMessage(message string, receiver pb.BroadcastMessageRequest_MessageType, entUser *ent.User, waitc chan BroadcastMessageResult) BroadcastMessageResult
}

func chunks(text string, limit int) []string {
	if len(text) <= limit {
		return []string{text}
	}
	var result []string
	for len(text) > 0 {
		if len(text) < limit {
			limit = len(text)
		}
		result = append(result, text[:limit])
		text = text[limit:]
	}
	return result
}

type contractNotifier struct {
	telegramNotifier      *telegramNotifier
	discordNotifier       *DiscordNotifier
	telegramChatManager   database.ITelegramChatManager
	discordChannelManager database.IDiscordChannelManager
}

func NewContractNotifier(
	managers *database.DbManagers,
	telegramBotToken string,
	telegramEndpoint string,
	discordBotToken string,
) ContractNotifier {
	return &contractNotifier{
		telegramNotifier:      newTelegramNotifier(managers, telegramBotToken, telegramEndpoint),
		discordNotifier:       newDiscordNotifier(managers, discordBotToken),
		telegramChatManager:   managers.TelegramChatManager,
		discordChannelManager: managers.DiscordChannelManager,
	}
}

func (n *contractNotifier) Notify(entContract *ent.Contract, entProp *ent.ContractProposal) {
	log.Sugar.Infof("Notifying for proposal %v on contract %v", entProp.ProposalID, entContract.Name)

	tgIds := n.telegramChatManager.QuerySubscribedIds(entContract.QueryTelegramChats())
	n.telegramNotifier.notify(tgIds, entContract.Name, entProp.ProposalID, entProp.Title, entProp.Description)

	discordIds := n.discordChannelManager.QuerySubscribedIds(entContract.QueryDiscordChannels())
	n.discordNotifier.notify(discordIds, entContract.Name, entProp.ProposalID, entProp.Title, entProp.Description)
}

type chainNotifier struct {
	telegramNotifier      *telegramNotifier
	discordNotifier       *DiscordNotifier
	telegramChatManager   database.ITelegramChatManager
	discordChannelManager database.IDiscordChannelManager
}

func NewChainNotifier(managers *database.DbManagers, telegramBotToken string, telegramEndpoint string, discordBotToken string) ChainNotifier {
	return &chainNotifier{
		telegramNotifier:      newTelegramNotifier(managers, telegramBotToken, telegramEndpoint),
		discordNotifier:       newDiscordNotifier(managers, discordBotToken),
		telegramChatManager:   managers.TelegramChatManager,
		discordChannelManager: managers.DiscordChannelManager,
	}
}

func (n *chainNotifier) Notify(entChain *ent.Chain, entProp *ent.ChainProposal) {
	log.Sugar.Infof("Notifying for proposal %v on chain %v", entProp.ProposalID, entChain.PrettyName)

	tgIds := n.telegramChatManager.QuerySubscribedIds(entChain.QueryTelegramChats())
	n.telegramNotifier.notify(tgIds, entChain.PrettyName, entProp.ProposalID, entProp.Title, entProp.Description)

	discordIds := n.discordChannelManager.QuerySubscribedIds(entChain.QueryDiscordChannels())
	n.discordNotifier.notify(discordIds, entChain.PrettyName, entProp.ProposalID, entProp.Title, entProp.Description)
}

func singularOrPluralText(count int, singular string, plural string) string {
	if count > 1 {
		return plural
	}
	return singular
}

func createRemainingTimeText(remainingTime time.Duration) string {
	if remainingTime > 0 {
		var dayText = "day"
		var hourText = "hour"
		var minuteText = "minute"
		if remainingTime.Hours() > 24 {
			var remainingDays = int(remainingTime.Hours() / 24)
			var remainingHours = int(remainingTime.Hours()) % 24
			var remainingMinutes = int(remainingTime.Minutes()) % 60
			dayText = singularOrPluralText(remainingDays, "day", "days")
			hourText = singularOrPluralText(remainingHours, "hour", "hours")
			minuteText = singularOrPluralText(remainingMinutes, "minute", "minutes")
			return fmt.Sprintf("%d %s %d %s %d %s", remainingDays, dayText, remainingHours, hourText, remainingMinutes, minuteText)
		} else if remainingTime.Hours() > 1 {
			var remainingHours = int(remainingTime.Hours()) % 24
			var remainingMinutes = int(remainingTime.Minutes()) % 60
			hourText = singularOrPluralText(remainingHours, "hour", "hours")
			minuteText = singularOrPluralText(remainingMinutes, "minute", "minutes")
			return fmt.Sprintf("%d %s %d %s", remainingHours, hourText, remainingMinutes, minuteText)
		} else {
			var remainingMinutes = int(remainingTime.Minutes()) % 60
			minuteText = singularOrPluralText(remainingMinutes, "minute", "minutes")
			return fmt.Sprintf("%d %s", remainingMinutes, minuteText)
		}
	}
	return ""
}

func (n *chainNotifier) SendVoteReminder(data database.AddressTrackerWithChainProposal) {
	var remainingTime = data.ChainProposal.VotingEndTime.Sub(time.Now())
	var remainingTimeText = createRemainingTimeText(remainingTime)

	if data.AddressTracker.Edges.TelegramChat != nil {
		n.telegramNotifier.sendVoteReminder(
			data.AddressTracker.Edges.TelegramChat,
			data.ChainProposal.Edges.Chain.PrettyName,
			data.ChainProposal.ProposalID,
			data.ChainProposal.Title,
			remainingTimeText,
		)
	}
	if data.AddressTracker.Edges.DiscordChannel != nil {
		n.discordNotifier.sendVoteReminder(
			data.AddressTracker.Edges.DiscordChannel,
			data.ChainProposal.Edges.Chain.PrettyName,
			data.ChainProposal.ProposalID,
			data.ChainProposal.Title,
			remainingTimeText,
		)
	}
}

type generalNotifier struct {
	telegramNotifier      *telegramNotifier
	discordNotifier       *DiscordNotifier
	telegramChatManager   database.ITelegramChatManager
	discordChannelManager database.IDiscordChannelManager
}

func NewGeneralNotifier(managers *database.DbManagers, telegramBotToken string, telegramEndpoint string, discordBotToken string) GeneralNotifier {
	return &generalNotifier{
		telegramNotifier:      newTelegramNotifier(managers, telegramBotToken, telegramEndpoint),
		discordNotifier:       newDiscordNotifier(managers, discordBotToken),
		telegramChatManager:   managers.TelegramChatManager,
		discordChannelManager: managers.DiscordChannelManager,
	}
}

type BroadcastMessageResult struct {
	ChatCnt        int
	SingleChatName string
	ErrorCnt       int
	IsSending      bool
	Error          error
}

func (n *generalNotifier) BroadcastMessage(message string, receiver pb.BroadcastMessageRequest_MessageType, entUser *ent.User, waitc chan BroadcastMessageResult) BroadcastMessageResult {
	switch receiver {
	case pb.BroadcastMessageRequest_TELEGRAM:
		tgIds := n.telegramChatManager.QueryAllIds()
		waitc <- BroadcastMessageResult{ChatCnt: len(tgIds), IsSending: true}
		errCnt := n.telegramNotifier.broadcastMessage(tgIds, message)
		return BroadcastMessageResult{ChatCnt: len(tgIds) - errCnt, ErrorCnt: errCnt}
	case pb.BroadcastMessageRequest_DISCORD:
		discordIds := n.discordChannelManager.QueryAllIds()
		waitc <- BroadcastMessageResult{ChatCnt: len(discordIds), IsSending: true}
		errCnt := n.discordNotifier.broadcastMessage(discordIds, message)
		return BroadcastMessageResult{ChatCnt: len(discordIds) - errCnt, ErrorCnt: errCnt}
	case pb.BroadcastMessageRequest_TELEGRAM_TEST:
		ids := n.telegramChatManager.QuerySubscribedIds(entUser.QueryTelegramChats())
		if len(ids) == 0 {
			return BroadcastMessageResult{Error: errors.New("no telegram chats found")}
		}
		errCnt := n.telegramNotifier.broadcastMessage(ids[:1], message)
		return BroadcastMessageResult{ChatCnt: 1 - errCnt, ErrorCnt: errCnt, SingleChatName: ids[0].Name}
	case pb.BroadcastMessageRequest_DISCORD_TEST:
		ids := n.discordChannelManager.QuerySubscribedIds(entUser.QueryDiscordChannels())
		if len(ids) == 0 {
			return BroadcastMessageResult{Error: errors.New("no discord channels found")}
		}
		errCnt := n.discordNotifier.broadcastMessage(ids[:1], message)
		return BroadcastMessageResult{ChatCnt: 1 - errCnt, ErrorCnt: errCnt, SingleChatName: ids[0].Name}
	}
	return BroadcastMessageResult{Error: errors.New("unknown receiver")}
}
