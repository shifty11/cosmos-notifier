-- reverse: create "discord_channel_users" table
DROP TABLE "discord_channel_users";
-- reverse: modify "discord_channels" table
ALTER TABLE "discord_channels" ADD COLUMN "discord_channel_user" bigint NULL;
-- reverse: create "telegram_chat_users" table
DROP TABLE "telegram_chat_users";
-- reverse: modify "telegram_chats" table
ALTER TABLE "telegram_chats" ADD COLUMN "telegram_chat_user" bigint NULL;
