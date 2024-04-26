-- +migrate Up
ALTER TABLE "permissions"
ADD CONSTRAINT "permissions_group_id_fk"
FOREIGN KEY("group_id") REFERENCES "users_groups"("id")
ON UPDATE NO ACTION ON DELETE CASCADE;
-- +migrate Down
ALTER TABLE "tokens"
DROP CONSTRAINT "permissions_group_id_fk";
