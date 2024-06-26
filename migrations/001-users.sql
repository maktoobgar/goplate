-- +migrate Up
CREATE TABLE "users" (
	"id" serial NOT NULL UNIQUE,
	"phone_number" varchar(16) NOT NULL,
	"phone_number_verified" BOOLEAN NOT NULL DEFAULT FALSE,
	"email" varchar(64),
	"email_verified" BOOLEAN NOT NULL DEFAULT FALSE,
	"password" varchar(256) NOT NULL,
	"avatar" varchar(256),
	"first_name" varchar(128),
	"last_name" varchar(128),
	"display_name" varchar(128) NOT NULL,
	-- 0 notdefined 1 male 2 female
	"gender" int NOT NULL DEFAULT 0,
	"is_active" boolean NOT NULL DEFAULT TRUE,
	"registered" boolean NOT NULL DEFAULT FALSE,
	"deactivation_reason" varchar(256),
	"is_admin" boolean NOT NULL DEFAULT FALSE,
	-- These can be inside params:
	-- 
	-- "otp_remaining_attempts" int NOT NULL DEFAULT 0,
	-- "otp_code" int,
	-- "otp_due_date" TIMESTAMPTZ,
	"params" varchar(512),
	"is_superuser" boolean NOT NULL DEFAULT FALSE,
	"created_at" TIMESTAMPTZ NOT NULL,
	PRIMARY KEY("id")
);
-- +migrate Down
DROP TABLE "users";
