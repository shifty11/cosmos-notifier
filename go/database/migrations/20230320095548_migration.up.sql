-- modify "address_trackers" table
ALTER TABLE "address_trackers" ADD COLUMN "notification_interval" bigint NOT NULL;
