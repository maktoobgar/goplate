CREATE TABLE "users" (
	"id" serial NOT NULL,
	"phone_number" varchar(16) NOT NULL,
	"email" varchar(64),
	"password" varchar(256) NOT NULL,
	"profile" varchar(256),
	"first_name" varchar(128),
	"last_name" varchar(128),
	"display_name" varchar(128) NOT NULL,
	-- 0 notdefined 1 male 2 female
	"gender" int NOT NULL DEFAULT 0,
	"is_active" boolean NOT NULL DEFAULT TRUE,
	"registered" boolean NOT NULL DEFAULT FALSE,
	"deactivation_reason" varchar(256),
	"is_admin" boolean NOT NULL DEFAULT FALSE,
	"otp_remaining_attempts" int NOT NULL DEFAULT 0,
	"otp_code" int,
	"otp_due_date" TIMESTAMPTZ,
	"is_superuser" boolean NOT NULL DEFAULT FALSE,
	"created_at" TIMESTAMPTZ NOT NULL,
	PRIMARY KEY("id")
);

CREATE TABLE "tokens" (
	"id" serial NOT NULL,
	"token" varchar(256) NOT NULL,
	"is_refresh_token" boolean NOT NULL DEFAULT FALSE,
	"user_id" int NOT NULL,
	"expires_at" TIMESTAMPTZ NOT NULL,
	"created_at" TIMESTAMPTZ,
	PRIMARY KEY("id")
);

CREATE TABLE "permissions" (
	"id" serial NOT NULL,
	"permission_id" int NOT NULL,
	"name" varchar(128) NOT NULL,
	"user_id" int,
	"group_id" int,
	"created_at" TIMESTAMPTZ NOT NULL,
	PRIMARY KEY("id")
);

CREATE TABLE "groups" (
	"id" serial NOT NULL,
	"name" varchar(255) NOT NULL,
	"created_at" TIMESTAMPTZ NOT NULL,
	PRIMARY KEY("id")
);

CREATE TABLE "users_groups" (
	"id" serial NOT NULL,
	"user_id" int NOT NULL,
	"group_id" int NOT NULL,
	"created_at" TIMESTAMPTZ,
	PRIMARY KEY("id")
);
