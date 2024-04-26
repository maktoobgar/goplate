-- +migrate Up
CREATE TABLE "groups" (
	"id" serial NOT NULL,
	"name" varchar(255) NOT NULL,
	"created_at" TIMESTAMPTZ NOT NULL,
	PRIMARY KEY("id")
);
-- +migrate Down
DROP TABLE "groups";
