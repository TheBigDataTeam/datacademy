CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "surname" varchar NOT NULL,
  "password" varchar NOT NULL,
  "isAdmin" boolean DEFAULT false,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "version" integer NOT NULL DEFAULT 1
);

CREATE TABLE "courses" (
  "id" varchar PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "theme" varchar,
  "author" varchar NOT NULL,
  "author_id" varchar NOT NULL,
  "techstack" varchar NOT NULL,
  "syllabus" varchar NOT NULL,
  "duration" varchar NOT NULL,
  "beneficiars" varchar NOT NULL,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "version" integer NOT NULL DEFAULT 1
);

CREATE TABLE "authors" (
  "id" varchar PRIMARY KEY,
  "course_id" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "fullname" varchar NOT NULL,
  "twitter_acc" varchar,
  "facebook_acc" varchar,
  "instagram_acc" varchar,
  "location" varchar,
  "bio" varchar NOT NULL,
  "shortdescription" varchar,
  "speciality" varchar,
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

ALTER TABLE "courses" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "authors" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "subscriptions" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");

ALTER TABLE "subscriptions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
