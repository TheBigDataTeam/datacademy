CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "surname" varchar NOT NULL,
  "password" bytea NOT NULL,
  "isAdmin" boolean DEFAULT false,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "version" integer NOT NULL DEFAULT 1
);

CREATE TABLE "courses" (
  "id" varchar PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "author" varchar NOT NULL,
  "techstack" varchar NOT NULL,
  "modulequantity" varchar NOT NULL,
  "workshopquantity" varchar NOT NULL,
  "diploma" varchar,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "version" integer NOT NULL DEFAULT 1
);

CREATE TABLE "authors" (
  "id" varchar PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "fullname" varchar NOT NULL,
  "twitter" varchar NOT NULL,
  "facebook" varchar NOT NULL,
  "instagram" varchar NOT NULL,
  "location" varchar NOT NULL,
  "bio" varchar NOT NULL,
  "shortdescription" varchar NOT NULL,
  "features" varchar NOT NULL,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "version" integer NOT NULL DEFAULT 1
);

CREATE TABLE "subscriptions" (
  "id" varchar PRIMARY KEY,
  "course_id" varchar NOT NULL,
  "user_id" varchar NOT NULL
);

CREATE TABLE "sessions" (
  "id" varchar PRIMARY KEY,
  "user_id" varchar NOT NULL
);

CREATE TABLE "modules" (
  "id" varchar PRIMARY KEY,
  "course_id" varchar NOT NULL,
  "title" varchar NOT NULL,
  "body" varchar NOT NULL,
  "badge" varchar,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "version" integer NOT NULL DEFAULT 1
);

CREATE TABLE "lessons" (
  "id" varchar PRIMARY KEY,
  "module_id" varchar NOT NULL,
  "title" varchar NOT NULL,
  "body" varchar NOT NULL,
  "video" varchar,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "version" integer NOT NULL DEFAULT 1
);

ALTER TABLE "subscriptions" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "subscriptions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "modules" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "lessons" ADD FOREIGN KEY ("module_id") REFERENCES "modules" ("id");
