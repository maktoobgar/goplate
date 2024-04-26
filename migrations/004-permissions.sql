-- +migrate Up
CREATE TABLE "permissions" (
	"id" serial NOT NULL,
	"permission_id" int NOT NULL,
	"name" varchar(128) NOT NULL,
	"user_id" int,
	"group_id" int,
	"created_at" TIMESTAMPTZ NOT NULL,
	PRIMARY KEY("id")
);
-- +migrate Down
DROP TABLE "permissions";
