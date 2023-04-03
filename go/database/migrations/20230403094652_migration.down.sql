-- reverse: create index "validator_operator_address_chain_validators" to table: "validators"
DROP INDEX "validator_operator_address_chain_validators";
-- reverse: drop index "validators_operator_address_key" from table: "validators"
CREATE UNIQUE INDEX "validators_operator_address_key" ON "validators" ("operator_address");
