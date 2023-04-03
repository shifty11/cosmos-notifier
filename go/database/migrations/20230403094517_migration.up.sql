-- drop index "validators_address_key" from table: "validators"
DROP INDEX "validators_address_key";
-- create index "validator_address_chain_validators" to table: "validators"
CREATE UNIQUE INDEX "validator_address_chain_validators" ON "validators" ("address", "chain_validators");
