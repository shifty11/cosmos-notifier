-- drop index "chain_name" from table: "chains"
DROP INDEX "chain_name";
-- drop index "chains_name_key" from table: "chains"
DROP INDEX "chains_name_key";
-- create index "chain_name" to table: "chains"
CREATE INDEX "chain_name" ON "chains" ("name");
