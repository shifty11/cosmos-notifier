-- reverse: create index "addresstracker_address_chain_a_3e2f36cdcb6f4a3f04a660254f75606e" to table: "address_trackers"
DROP INDEX "addresstracker_address_chain_a_3e2f36cdcb6f4a3f04a660254f75606e";
-- reverse: modify "address_trackers" table
ALTER TABLE "address_trackers" DROP CONSTRAINT "address_trackers_validators_address_trackers", DROP COLUMN "validator_address_trackers";
-- reverse: drop index "addresstracker_address_chain_a_3e2f36cdcb6f4a3f04a660254f75606e" from table: "address_trackers"
CREATE UNIQUE INDEX "addresstracker_address_chain_a_3e2f36cdcb6f4a3f04a660254f75606e" ON "address_trackers" ("address", "chain_address_trackers", "discord_channel_address_trackers", "telegram_chat_address_trackers");
-- reverse: create index "validator_moniker" to table: "validators"
DROP INDEX "validator_moniker";
-- reverse: create index "validator_address" to table: "validators"
DROP INDEX "validator_address";
-- reverse: create "validators" table
DROP TABLE "validators";
-- reverse: modify "chains" table
ALTER TABLE "chains" ALTER COLUMN "bech32_prefix" SET DEFAULT '';
