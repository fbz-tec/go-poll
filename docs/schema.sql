-- SQL dump generated using DBML (dbml.dbdiagram.io)
-- Database: PostgreSQL
-- Generated at: 2025-02-06T12:19:55.848Z

CREATE TABLE "polls" (
  "poll_id" bigserial PRIMARY KEY,
  "question" text NOT NULL,
  "owner" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "options" (
  "option_id" bigserial PRIMARY KEY,
  "option_value" text NOT NULL,
  "poll_id" bigint NOT NULL
);

CREATE TABLE "votes" (
  "vote_id" bigserial PRIMARY KEY,
  "option_id" bigint NOT NULL,
  "voter" varchar NOT NULL
);

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "role" varchar NOT NULL DEFAULT 'voter',
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "is_email_verified" bool NOT NULL DEFAULT false,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "options" ADD FOREIGN KEY ("poll_id") REFERENCES "polls" ("poll_id");

ALTER TABLE "polls" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "votes" ADD FOREIGN KEY ("voter") REFERENCES "users" ("username");
