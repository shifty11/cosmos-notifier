-- modify "validators" table
ALTER TABLE "validators" ADD COLUMN "first_inactive_time" timestamptz NULL;
