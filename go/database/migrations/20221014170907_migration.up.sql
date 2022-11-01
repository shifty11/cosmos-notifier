-- create "chains" table
CREATE TABLE "chains" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "create_time" timestamptz NOT NULL, "update_time" timestamptz NOT NULL, "chain_id" character varying NOT NULL, "name" character varying NOT NULL, "pretty_name" character varying NOT NULL, "is_enabled" boolean NOT NULL DEFAULT false, "image_url" character varying NOT NULL, "thumbnail_url" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- create index "chains_chain_id_key" to table: "chains"
CREATE UNIQUE INDEX "chains_chain_id_key" ON "chains" ("chain_id");
-- create index "chains_name_key" to table: "chains"
CREATE UNIQUE INDEX "chains_name_key" ON "chains" ("name");
-- create index "chains_pretty_name_key" to table: "chains"
CREATE UNIQUE INDEX "chains_pretty_name_key" ON "chains" ("pretty_name");
-- create index "chain_name" to table: "chains"
CREATE UNIQUE INDEX "chain_name" ON "chains" ("name");
-- create "chain_proposals" table
CREATE TABLE "chain_proposals" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "create_time" timestamptz NOT NULL, "update_time" timestamptz NOT NULL, "proposal_id" bigint NOT NULL, "title" character varying NOT NULL, "description" character varying NOT NULL, "voting_start_time" timestamptz NOT NULL, "voting_end_time" timestamptz NOT NULL, "status" character varying NOT NULL, "chain_chain_proposals" bigint NULL, PRIMARY KEY ("id"), CONSTRAINT "chain_proposals_chains_chain_proposals" FOREIGN KEY ("chain_chain_proposals") REFERENCES "chains" ("id") ON DELETE CASCADE);
-- create "discord_channel_chains" table
CREATE TABLE "discord_channel_chains" ("discord_channel_id" bigint NOT NULL, "chain_id" bigint NOT NULL, PRIMARY KEY ("discord_channel_id", "chain_id"), CONSTRAINT "discord_channel_chains_discord_channel_id" FOREIGN KEY ("discord_channel_id") REFERENCES "discord_channels" ("id") ON DELETE CASCADE, CONSTRAINT "discord_channel_chains_chain_id" FOREIGN KEY ("chain_id") REFERENCES "chains" ("id") ON DELETE CASCADE);
-- create "telegram_chat_chains" table
CREATE TABLE "telegram_chat_chains" ("telegram_chat_id" bigint NOT NULL, "chain_id" bigint NOT NULL, PRIMARY KEY ("telegram_chat_id", "chain_id"), CONSTRAINT "telegram_chat_chains_telegram_chat_id" FOREIGN KEY ("telegram_chat_id") REFERENCES "telegram_chats" ("id") ON DELETE CASCADE, CONSTRAINT "telegram_chat_chains_chain_id" FOREIGN KEY ("chain_id") REFERENCES "chains" ("id") ON DELETE CASCADE);