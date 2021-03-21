CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "surname" varchar NOT NULL,
  "subscribed_to" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "isAdmin" boolean DEFAULT false
);

CREATE TABLE "courses" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "author_id" bigint NOT NULL,
  "theme" varchar NOT NULL,
  "difficulty" varchar NOT NULL,
  "duration" varchar NOT NULL,
  "tech_stack" varchar NOT NULL,
  "created_on" timestamptz NOT NULL DEFAULT (now()),
  "updated_on" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "authors" (
  "id" bigserial PRIMARY KEY,
  "course_id" bigint NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "surname" varchar NOT NULL,
  "speciality" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "users" ADD FOREIGN KEY ("subscribed_to") REFERENCES "courses" ("id");

ALTER TABLE "courses" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "authors" ADD FOREIGN KEY ("course_id") REFERENCES "courses" ("id");
