-- +migrate Up
CREATE TABLE "tokens" (
	"id" serial NOT NULL,
	"user_id" int NOT NULL,
	"created_at" TIMESTAMPTZ NOT NULL,
	PRIMARY KEY("id")
);
-- +migrate Down
DROP TABLE "tokens";
