-- reverse: create "address_tracker_chain_proposals" table
DROP TABLE "address_tracker_chain_proposals";
-- reverse: create index "addresstracker_address_chain_a_3e2f36cdcb6f4a3f04a660254f75606e" to table: "address_trackers"
DROP INDEX "addresstracker_address_chain_a_3e2f36cdcb6f4a3f04a660254f75606e";
-- reverse: create "address_trackers" table
DROP TABLE "address_trackers";
-- reverse: modify "chains" table
ALTER TABLE "chains" DROP COLUMN "bech32_prefix";
