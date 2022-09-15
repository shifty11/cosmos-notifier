-- modify "telegram_chats" table
ALTER TABLE "telegram_chats" DROP COLUMN "telegram_chat_user";
-- create "telegram_chat_users" table
CREATE TABLE "telegram_chat_users" ("telegram_chat_id" bigint NOT NULL, "user_id" bigint NOT NULL, PRIMARY KEY ("telegram_chat_id", "user_id"), CONSTRAINT "telegram_chat_users_telegram_chat_id" FOREIGN KEY ("telegram_chat_id") REFERENCES "telegram_chats" ("id") ON DELETE CASCADE, CONSTRAINT "telegram_chat_users_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE);
-- modify "discord_channels" table
ALTER TABLE "discord_channels" DROP COLUMN "discord_channel_user";
-- create "discord_channel_users" table
CREATE TABLE "discord_channel_users" ("discord_channel_id" bigint NOT NULL, "user_id" bigint NOT NULL, PRIMARY KEY ("discord_channel_id", "user_id"), CONSTRAINT "discord_channel_users_discord_channel_id" FOREIGN KEY ("discord_channel_id") REFERENCES "discord_channels" ("id") ON DELETE CASCADE, CONSTRAINT "discord_channel_users_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE);
