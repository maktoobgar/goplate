-- +migrate Up
CREATE TABLE "users_groups" (
	"id" serial NOT NULL,
	"user_id" int NOT NULL,
	"group_id" int NOT NULL,
	"created_at" TIMESTAMPTZ,
	PRIMARY KEY("id")
);
-- +migrate Down
DROP TABLE "users_groups";
