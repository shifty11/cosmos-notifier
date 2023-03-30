-- modify "validators" table
ALTER TABLE "validators" ADD COLUMN "operator_address" character varying NOT NULL;
-- create index "validator_moniker_operator_address_chain_validators" to table: "validators"
CREATE UNIQUE INDEX "validator_moniker_operator_address_chain_validators" ON "validators" ("moniker", "operator_address", "chain_validators");
-- create index "validator_operator_address" to table: "validators"
CREATE INDEX "validator_operator_address" ON "validators" ("operator_address");
-- create index "validators_operator_address_key" to table: "validators"
CREATE UNIQUE INDEX "validators_operator_address_key" ON "validators" ("operator_address");
