-- Create "users" table
CREATE TABLE "users" (
  "id" uuid NOT NULL,
  "deleted_at" timestamptz NULL,
  "email" character varying NOT NULL,
  "password" character varying NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- Create "notes" table
CREATE TABLE "notes" (
  "id" uuid NOT NULL,
  "deleted_at" timestamptz NULL,
  "share_key" character varying NOT NULL,
  "title" character varying NULL,
  "content" text NOT NULL DEFAULT '',
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "user_id" uuid NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "notes_users_Note" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL
);
-- Create index "notes_share_key_key" to table: "notes"
CREATE UNIQUE INDEX "notes_share_key_key" ON "notes" ("share_key");
