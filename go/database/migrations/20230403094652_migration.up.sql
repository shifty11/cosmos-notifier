-- drop index "validators_operator_address_key" from table: "validators"
DROP INDEX "validators_operator_address_key";
-- create index "validator_operator_address_chain_validators" to table: "validators"
CREATE UNIQUE INDEX "validator_operator_address_chain_validators" ON "validators" ("operator_address", "chain_validators");
