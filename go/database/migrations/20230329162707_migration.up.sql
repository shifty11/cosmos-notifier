-- create index "validator_address_moniker" to table: "validators"
CREATE UNIQUE INDEX "validator_address_moniker" ON "validators" ("address", "moniker");
