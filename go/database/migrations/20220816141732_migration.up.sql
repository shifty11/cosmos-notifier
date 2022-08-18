-- modify "contracts" table
ALTER TABLE "contracts" DROP COLUMN "discord_channel_chains", DROP COLUMN "telegram_chat_contracts";
-- create "discord_channel_contracts" table
CREATE TABLE "discord_channel_contracts" ("discord_channel_id" bigint NOT NULL, "contract_id" bigint NOT NULL, PRIMARY KEY ("discord_channel_id", "contract_id"), CONSTRAINT "discord_channel_contracts_discord_channel_id" FOREIGN KEY ("discord_channel_id") REFERENCES "discord_channels" ("id") ON DELETE CASCADE, CONSTRAINT "discord_channel_contracts_contract_id" FOREIGN KEY ("contract_id") REFERENCES "contracts" ("id") ON DELETE CASCADE);
-- create "telegram_chat_contracts" table
CREATE TABLE "telegram_chat_contracts" ("telegram_chat_id" bigint NOT NULL, "contract_id" bigint NOT NULL, PRIMARY KEY ("telegram_chat_id", "contract_id"), CONSTRAINT "telegram_chat_contracts_telegram_chat_id" FOREIGN KEY ("telegram_chat_id") REFERENCES "telegram_chats" ("id") ON DELETE CASCADE, CONSTRAINT "telegram_chat_contracts_contract_id" FOREIGN KEY ("contract_id") REFERENCES "contracts" ("id") ON DELETE CASCADE);
