-- +migrate Up
ALTER TABLE "permissions"
ADD CONSTRAINT "permissions_user_id_fk"
FOREIGN KEY("user_id") REFERENCES "users"("id")
ON UPDATE NO ACTION ON DELETE CASCADE;
-- +migrate Down
ALTER TABLE "tokens"
DROP CONSTRAINT "permissions_user_id_fk";
