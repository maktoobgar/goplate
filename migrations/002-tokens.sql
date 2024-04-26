-- +migrate Up
CREATE TABLE "tokens" (
	"id" serial NOT NULL,
	"token" varchar(256) NOT NULL,
	"is_refresh_token" boolean NOT NULL DEFAULT FALSE,
	"user_id" int NOT NULL,
	"expires_at" TIMESTAMPTZ NOT NULL,
	"created_at" TIMESTAMPTZ,
	PRIMARY KEY("id")
);
-- +migrate Down
DROP TABLE "tokens";
