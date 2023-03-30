-- drop index "validator_moniker_chain_validators" from table: "validators"
DROP INDEX "validator_moniker_chain_validators";
-- create index "validator_moniker_address_chain_validators" to table: "validators"
CREATE UNIQUE INDEX "validator_moniker_address_chain_validators" ON "validators" ("moniker", "address", "chain_validators");
