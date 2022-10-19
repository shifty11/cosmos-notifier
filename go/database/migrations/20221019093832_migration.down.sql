-- reverse: rename "proposals" table
ALTER TABLE "contract_proposals" DROP CONSTRAINT "contract_proposals_contracts_proposals", ADD CONSTRAINT "proposals_contracts_proposals" FOREIGN KEY ("contract_proposals") REFERENCES "public"."contracts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
ALTER TABLE "contract_proposals" RENAME TO "proposals";
