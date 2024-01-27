CREATE TABLE IF NOT EXISTS "user"(
    "id" SERIAL,
    "guid" UUID NOT NULL,
    "username" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "password" TEXT NOT NULL,
    "created_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "user" ADD PRIMARY KEY("id");
ALTER TABLE "user" ADD CONSTRAINT "user_guid_unique" UNIQUE("guid");
ALTER TABLE "user" ADD CONSTRAINT "user_email_unique" UNIQUE("email");
ALTER TABLE "user" ADD CONSTRAINT "user_phone_unique" UNIQUE("phone");    