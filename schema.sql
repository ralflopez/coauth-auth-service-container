CREATE TYPE Role AS ENUM (
  'Member',
  'Admin',
  'Guest'
);

DROP TABLE IF EXISTS "users";
CREATE TABLE "public"."users" (
    "id" uuid NOT NULL,
    "name" character varying(255) NOT NULL,
    "email" character varying(255) NOT NULL,
    "password" text NOT NULL,
    "role" role DEFAULT Member NOT NULL,
    "created_at" timestamptz DEFAULT now(),
    "updated_at" timestamptz DEFAULT now()
) WITH (oids = false);
