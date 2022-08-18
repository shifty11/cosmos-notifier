-- reverse: create "telegram_chat_contracts" table
DROP TABLE "telegram_chat_contracts";
-- reverse: create "discord_channel_contracts" table
DROP TABLE "discord_channel_contracts";
-- reverse: modify "contracts" table
ALTER TABLE "contracts" ADD COLUMN "telegram_chat_contracts" bigint NULL, ADD COLUMN "discord_channel_chains" bigint NULL;
