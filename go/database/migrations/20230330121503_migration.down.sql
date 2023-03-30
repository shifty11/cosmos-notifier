-- reverse: create index "validators_operator_address_key" to table: "validators"
DROP INDEX "validators_operator_address_key";
-- reverse: create index "validator_operator_address" to table: "validators"
DROP INDEX "validator_operator_address";
-- reverse: create index "validator_moniker_operator_address_chain_validators" to table: "validators"
DROP INDEX "validator_moniker_operator_address_chain_validators";
-- reverse: modify "validators" table
ALTER TABLE "validators" DROP COLUMN "operator_address";
