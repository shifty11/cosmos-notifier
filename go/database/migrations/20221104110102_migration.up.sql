-- modify "contracts" table
ALTER TABLE "contracts" DROP COLUMN "get_config_json", DROP COLUMN "get_proposals_json", ADD COLUMN "config_version" character varying NOT NULL DEFAULT 'unknown';
