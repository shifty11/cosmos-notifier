package notifier

import (
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/log"
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
	telegramNotifier      *TelegramNotifier
	discordNotifier       *DiscordNotifier
	telegramChatManager   database.ITelegramChatManager
	discordChannelManager database.IDiscordChannelManager
}

func NewContractNotifier(managers *database.DbManagers, telegramBotToken string, telegramEndpoint string, discordBotToken string) *ContractNotifier {
	return &ContractNotifier{
		telegramNotifier:      NewTelegramNotifier(managers, telegramBotToken, telegramEndpoint),
		discordNotifier:       NewDiscordNotifier(managers, discordBotToken),
		telegramChatManager:   managers.TelegramChatManager,
		discordChannelManager: managers.DiscordChannelManager,
	}
}

func (n *ContractNotifier) Notify(entContract *ent.Contract, entProp *ent.ContractProposal) {
	log.Sugar.Infof("Notifying for proposal %v on contract %v", entProp.ProposalID, entContract.Name)

	tgIds := n.telegramChatManager.GetSubscribedIds(entContract.QueryTelegramChats())
	n.telegramNotifier.Notify(tgIds, entContract.Name, entProp.ProposalID, entProp.Title, entProp.Description)

	discordIds := n.discordChannelManager.GetSubscribedIds(entContract.QueryDiscordChannels())
	n.discordNotifier.Notify(discordIds, entContract.Name, entProp.ProposalID, entProp.Title, entProp.Description)
}

type ChainNotifier struct {
	telegramNotifier      *TelegramNotifier
	discordNotifier       *DiscordNotifier
	telegramChatManager   database.ITelegramChatManager
	discordChannelManager database.IDiscordChannelManager
}

func NewChainNotifier(managers *database.DbManagers, telegramBotToken string, telegramEndpoint string, discordBotToken string) *ChainNotifier {
	return &ChainNotifier{
		telegramNotifier:      NewTelegramNotifier(managers, telegramBotToken, telegramEndpoint),
		discordNotifier:       NewDiscordNotifier(managers, discordBotToken),
		telegramChatManager:   managers.TelegramChatManager,
		discordChannelManager: managers.DiscordChannelManager,
	}
}

func (n *ChainNotifier) Notify(entChain *ent.Chain, entProp *ent.ChainProposal) {
	log.Sugar.Infof("Notifying for proposal %v on chain %v", entProp.ProposalID, entChain.PrettyName)

	tgIds := n.telegramChatManager.GetSubscribedIds(entChain.QueryTelegramChats())
	n.telegramNotifier.Notify(tgIds, entChain.Name, entProp.ProposalID, entProp.Title, entProp.Description)

	discordIds := n.discordChannelManager.GetSubscribedIds(entChain.QueryDiscordChannels())
	n.discordNotifier.Notify(discordIds, entChain.Name, entProp.ProposalID, entProp.Title, entProp.Description)
}
