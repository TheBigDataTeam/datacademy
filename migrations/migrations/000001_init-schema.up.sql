CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "fullname" varchar NOT NULL,
  "subscribe_to" bigint,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "isAdmin" boolean DEFAULT false
);

CREATE TABLE "courses" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "theme" varchar,
  "author" varchar NOT NULL,
  "author_id" bigint NOT NULL,
  "tech_stack" varchar NOT NULL,
  "syllabus" varchar NOT NULL,
  "duration" varchar NOT NULL,
  "beneficiars" varchar NOT NULL,
  "difficulty" varchar NOT NULL,
  "created_on" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "authors" (
  "id" bigserial PRIMARY KEY,
  "course_id" bigint NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "twitter" varchar NOT NULL,
  "facebook" varchar NOT NULL,
  "instagram" varchar NOT NULL,
  "location" varchar NOT NULL,
  "fullname" varchar NOT NULL,
  "bio" varchar NOT NULL,
  "shortdescription" varchar NOT NULL,
  "speciality" varchar,
  "features" varchar NOT NULL,
  "created_on" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "users" ADD FOREIGN KEY ("subscribe_to") REFERENCES "courses" ("id");

ALTER TABLE "courses" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "authors" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");
