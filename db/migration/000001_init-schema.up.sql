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