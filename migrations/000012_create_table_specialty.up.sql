CREATE TABLE IF NOT EXISTS "specialty"(
    "id" SERIAL,
    "name" VARCHAR(255) NOT NULL
);

ALTER TABLE "specialty" ADD PRIMARY KEY("id");

ALTER TABLE "doctor_specialty" ADD CONSTRAINT "doctor_specialty_specialty_id_foreign" FOREIGN KEY("specialty_id") REFERENCES "specialty"("id");