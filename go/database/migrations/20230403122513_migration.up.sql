-- modify "address_trackers" table
ALTER TABLE "address_trackers" DROP CONSTRAINT "address_trackers_validators_address_trackers", ADD CONSTRAINT "address_trackers_validators_address_trackers" FOREIGN KEY ("validator_address_trackers") REFERENCES "validators" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
