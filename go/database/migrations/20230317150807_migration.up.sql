-- modify "chains" table
ALTER TABLE "chains" ADD COLUMN "bech32_prefix" character varying NOT NULL DEFAULT '';
-- create "address_trackers" table
CREATE TABLE "address_trackers" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "create_time" timestamptz NOT NULL, "update_time" timestamptz NOT NULL, "address" character varying NOT NULL, "chain_address_trackers" bigint NOT NULL, "discord_channel_address_trackers" bigint NULL, "telegram_chat_address_trackers" bigint NULL, PRIMARY KEY ("id"), CONSTRAINT "address_trackers_chains_address_trackers" FOREIGN KEY ("chain_address_trackers") REFERENCES "chains" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "address_trackers_discord_channels_address_trackers" FOREIGN KEY ("discord_channel_address_trackers") REFERENCES "discord_channels" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "address_trackers_telegram_chats_address_trackers" FOREIGN KEY ("telegram_chat_address_trackers") REFERENCES "telegram_chats" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- create index "addresstracker_address_chain_a_3e2f36cdcb6f4a3f04a660254f75606e" to table: "address_trackers"
CREATE UNIQUE INDEX "addresstracker_address_chain_a_3e2f36cdcb6f4a3f04a660254f75606e" ON "address_trackers" ("address", "chain_address_trackers", "discord_channel_address_trackers", "telegram_chat_address_trackers");
-- create "address_tracker_chain_proposals" table
CREATE TABLE "address_tracker_chain_proposals" ("address_tracker_id" bigint NOT NULL, "chain_proposal_id" bigint NOT NULL, PRIMARY KEY ("address_tracker_id", "chain_proposal_id"), CONSTRAINT "address_tracker_chain_proposals_address_tracker_id" FOREIGN KEY ("address_tracker_id") REFERENCES "address_trackers" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "address_tracker_chain_proposals_chain_proposal_id" FOREIGN KEY ("chain_proposal_id") REFERENCES "chain_proposals" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);