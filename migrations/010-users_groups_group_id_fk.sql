-- +migrate Up
ALTER TABLE "users_groups"
ADD CONSTRAINT "users_groups_group_id_fk"
FOREIGN KEY("group_id") REFERENCES "groups"("id")
ON UPDATE NO ACTION ON DELETE CASCADE;
-- +migrate Down
ALTER TABLE "tokens"
DROP CONSTRAINT "users_groups_group_id_fk";
