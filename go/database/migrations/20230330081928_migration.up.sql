-- drop index "validator_address_moniker" from table: "validators"
DROP INDEX "validator_address_moniker";
-- create index "validator_moniker_chain_validators" to table: "validators"
CREATE UNIQUE INDEX "validator_moniker_chain_validators" ON "validators" ("moniker", "chain_validators");
