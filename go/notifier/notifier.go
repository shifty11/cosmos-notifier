package notifier

import (
	"errors"
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/log"
	pb "github.com/shifty11/cosmos-notifier/services/grpc/protobuf/go/admin_service"
)

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

type ContractNotifier struct {
	telegramNotifier      *telegramNotifier
	discordNotifier       *DiscordNotifier
	telegramChatManager   database.ITelegramChatManager
	discordChannelManager database.IDiscordChannelManager
}

func NewContractNotifier(managers *database.DbManagers, telegramBotToken string, telegramEndpoint string, discordBotToken string) *ContractNotifier {
	return &ContractNotifier{
		telegramNotifier:      newTelegramNotifier(managers, telegramBotToken, telegramEndpoint),
		discordNotifier:       newDiscordNotifier(managers, discordBotToken),
		telegramChatManager:   managers.TelegramChatManager,
		discordChannelManager: managers.DiscordChannelManager,
	}
}

func (n *ContractNotifier) Notify(entContract *ent.Contract, entProp *ent.ContractProposal) {
	log.Sugar.Infof("Notifying for proposal %v on contract %v", entProp.ProposalID, entContract.Name)

	tgIds := n.telegramChatManager.GetSubscribedIds(entContract.QueryTelegramChats())
	n.telegramNotifier.notify(tgIds, entContract.Name, entProp.ProposalID, entProp.Title, entProp.Description)

	discordIds := n.discordChannelManager.GetSubscribedIds(entContract.QueryDiscordChannels())
	n.discordNotifier.notify(discordIds, entContract.Name, entProp.ProposalID, entProp.Title, entProp.Description)
}

type ChainNotifier struct {
	telegramNotifier      *telegramNotifier
	discordNotifier       *DiscordNotifier
	telegramChatManager   database.ITelegramChatManager
	discordChannelManager database.IDiscordChannelManager
}

func NewChainNotifier(managers *database.DbManagers, telegramBotToken string, telegramEndpoint string, discordBotToken string) *ChainNotifier {
	return &ChainNotifier{
		telegramNotifier:      newTelegramNotifier(managers, telegramBotToken, telegramEndpoint),
		discordNotifier:       newDiscordNotifier(managers, discordBotToken),
		telegramChatManager:   managers.TelegramChatManager,
		discordChannelManager: managers.DiscordChannelManager,
	}
}

func (n *ChainNotifier) Notify(entChain *ent.Chain, entProp *ent.ChainProposal) {
	log.Sugar.Infof("Notifying for proposal %v on chain %v", entProp.ProposalID, entChain.PrettyName)

	tgIds := n.telegramChatManager.GetSubscribedIds(entChain.QueryTelegramChats())
	n.telegramNotifier.notify(tgIds, entChain.PrettyName, entProp.ProposalID, entProp.Title, entProp.Description)

	discordIds := n.discordChannelManager.GetSubscribedIds(entChain.QueryDiscordChannels())
	n.discordNotifier.notify(discordIds, entChain.PrettyName, entProp.ProposalID, entProp.Title, entProp.Description)
}

type GeneralNotifier struct {
	telegramNotifier      *telegramNotifier
	discordNotifier       *DiscordNotifier
	telegramChatManager   database.ITelegramChatManager
	discordChannelManager database.IDiscordChannelManager
}

func NewGeneralNotifier(managers *database.DbManagers, telegramBotToken string, telegramEndpoint string, discordBotToken string) *GeneralNotifier {
	return &GeneralNotifier{
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

func (n *GeneralNotifier) BroadcastMessage(message string, receiver pb.BroadcastMessageRequest_MessageType, entUser *ent.User, waitc chan BroadcastMessageResult) BroadcastMessageResult {
	switch receiver {
	case pb.BroadcastMessageRequest_TELEGRAM:
		tgIds := n.telegramChatManager.GetAllIds()
		waitc <- BroadcastMessageResult{ChatCnt: len(tgIds), IsSending: true}
		errCnt := n.telegramNotifier.broadcastMessage(tgIds, message)
		return BroadcastMessageResult{ChatCnt: len(tgIds) - errCnt, ErrorCnt: errCnt}
	case pb.BroadcastMessageRequest_DISCORD:
		discordIds := n.discordChannelManager.GetAllIds()
		waitc <- BroadcastMessageResult{ChatCnt: len(discordIds), IsSending: true}
		errCnt := n.discordNotifier.broadcastMessage(discordIds, message)
		return BroadcastMessageResult{ChatCnt: len(discordIds) - errCnt, ErrorCnt: errCnt}
	case pb.BroadcastMessageRequest_TELEGRAM_TEST:
		ids := n.telegramChatManager.GetSubscribedIds(entUser.QueryTelegramChats())
		if len(ids) == 0 {
			return BroadcastMessageResult{Error: errors.New("no telegram chats found")}
		}
		errCnt := n.telegramNotifier.broadcastMessage(ids[:1], message)
		return BroadcastMessageResult{ChatCnt: 1 - errCnt, ErrorCnt: errCnt, SingleChatName: ids[0].Name}
	case pb.BroadcastMessageRequest_DISCORD_TEST:
		ids := n.discordChannelManager.GetSubscribedIds(entUser.QueryDiscordChannels())
		if len(ids) == 0 {
			return BroadcastMessageResult{Error: errors.New("no discord channels found")}
		}
		errCnt := n.discordNotifier.broadcastMessage(ids[:1], message)
		return BroadcastMessageResult{ChatCnt: 1 - errCnt, ErrorCnt: errCnt, SingleChatName: ids[0].Name}
	}
	return BroadcastMessageResult{Error: errors.New("unknown receiver")}
}
