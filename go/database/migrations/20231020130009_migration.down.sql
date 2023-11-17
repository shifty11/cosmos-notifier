-- reverse: create index "chain_name" to table: "chains"
DROP INDEX "chain_name";
-- reverse: drop index "chains_name_key" from table: "chains"
CREATE UNIQUE INDEX "chains_name_key" ON "chains" ("name");
-- reverse: drop index "chain_name" from table: "chains"
CREATE UNIQUE INDEX "chain_name" ON "chains" ("name");
