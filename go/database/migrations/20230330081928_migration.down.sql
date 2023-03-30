-- reverse: create index "validator_moniker_chain_validators" to table: "validators"
DROP INDEX "validator_moniker_chain_validators";
-- reverse: drop index "validator_address_moniker" from table: "validators"
CREATE UNIQUE INDEX "validator_address_moniker" ON "validators" ("address", "moniker");
