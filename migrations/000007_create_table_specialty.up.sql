CREATE TABLE IF NOT EXISTS "specialty"(
    "id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "specialty" ADD PRIMARY KEY("id");

ALTER TABLE "specialty" ADD CONSTRAINT "specialty_name_unique" UNIQUE("name");

ALTER TABLE "doctor_specialty" ADD CONSTRAINT "doctorspecialty_specialty_id_foreign" FOREIGN KEY("specialty_id") REFERENCES "specialty"("id");