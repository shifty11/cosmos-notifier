-- reverse: drop index "chains_pretty_name_key" from table: "chains"
CREATE UNIQUE INDEX "chains_pretty_name_key" ON "chains" ("pretty_name");
