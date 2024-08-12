-- +migrate Up
ALTER TABLE "tokens"
ADD CONSTRAINT "tokens_user_id_fk"
FOREIGN KEY ("user_id") REFERENCES "users"("id")
ON DELETE CASCADE;
-- +migrate Down
ALTER TABLE "tokens"
DROP CONSTRAINT "tokens_user_id_fk";
