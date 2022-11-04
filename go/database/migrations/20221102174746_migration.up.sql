-- modify "contracts" table
ALTER TABLE "contracts" ADD COLUMN "rpc_endpoint" character varying NOT NULL DEFAULT 'https://rpc.cosmos.directory/juno', ADD COLUMN "get_config_json" character varying NOT NULL DEFAULT '', ADD COLUMN "get_proposals_json" character varying NOT NULL DEFAULT '';
