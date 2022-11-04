-- reverse: modify "contracts" table
ALTER TABLE "contracts" DROP COLUMN "config_version", ADD COLUMN "get_proposals_json" character varying NOT NULL DEFAULT '', ADD COLUMN "get_config_json" character varying NOT NULL DEFAULT '';
