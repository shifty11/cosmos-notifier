package types

type TgChatQueryResult struct {
	ChatId int64  `json:"chat_id"`
	Name   string `json:"name,omitempty"`
}

type DiscordChannelQueryResult struct {
	ChannelId int64  `json:"channel_id"`
	Name      string `json:"name,omitempty"`
}
