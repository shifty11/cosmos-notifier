-- reverse: modify "contracts" table
ALTER TABLE "contracts" DROP COLUMN "get_proposals_json", DROP COLUMN "get_config_json", DROP COLUMN "rpc_endpoint";
