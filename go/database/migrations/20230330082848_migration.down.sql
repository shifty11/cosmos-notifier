-- reverse: create index "validator_moniker_address_chain_validators" to table: "validators"
DROP INDEX "validator_moniker_address_chain_validators";
-- reverse: drop index "validator_moniker_chain_validators" from table: "validators"
CREATE UNIQUE INDEX "validator_moniker_chain_validators" ON "validators" ("moniker", "chain_validators");
