-- rename "proposals" table
ALTER TABLE "proposals" RENAME TO "contract_proposals";
ALTER TABLE "contract_proposals" DROP CONSTRAINT "proposals_contracts_proposals", ADD CONSTRAINT "contract_proposals_contracts_proposals" FOREIGN KEY ("contract_proposals") REFERENCES "contracts" ("id") ON DELETE CASCADE;
