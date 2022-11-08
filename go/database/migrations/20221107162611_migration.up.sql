-- modify "chains" table
ALTER TABLE "chains" ADD COLUMN "path" character varying NOT NULL DEFAULT '', ADD COLUMN "display" character varying NOT NULL DEFAULT '';
-- modify "contracts" table
ALTER TABLE "contracts" ADD COLUMN "get_proposals_query" character varying NOT NULL DEFAULT '{"list_proposals":{}}';
