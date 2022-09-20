-- modify "users" table
ALTER TABLE "users" ADD COLUMN "role" character varying NOT NULL DEFAULT 'user';
