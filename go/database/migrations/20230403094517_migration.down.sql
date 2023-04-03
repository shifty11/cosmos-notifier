-- reverse: create index "validator_address_chain_validators" to table: "validators"
DROP INDEX "validator_address_chain_validators";
-- reverse: drop index "validators_address_key" from table: "validators"
CREATE UNIQUE INDEX "validators_address_key" ON "validators" ("address");
