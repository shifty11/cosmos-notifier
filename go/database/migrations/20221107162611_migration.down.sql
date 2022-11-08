-- reverse: modify "contracts" table
ALTER TABLE "contracts" DROP COLUMN "get_proposals_query";
-- reverse: modify "chains" table
ALTER TABLE "chains" DROP COLUMN "display", DROP COLUMN "path";
