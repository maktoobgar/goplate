-- +migrate Up
ALTER TABLE "users_groups"
ADD CONSTRAINT "users_groups_user_id_fk"
FOREIGN KEY("user_id") REFERENCES "users"("id")
ON UPDATE NO ACTION ON DELETE CASCADE;
-- +migrate Down
ALTER TABLE "users_groups"
DROP CONSTRAINT "users_groups_user_id_fk";
