CREATE TABLE IF NOT EXISTS "doctor"(
    "doctor_id" BIGINT NOT NULL,
    "clinic_id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "surname" VARCHAR(255) NOT NULL,
    "birthday" DATE NOT NULL,
    "gender" VARCHAR(255) NOT NULL,
    "education" TEXT NOT NULL,
    "certificates" TEXT NOT NULL
);

CREATE INDEX "doctor_name_index" ON "doctor"("name");

CREATE INDEX "doctor_surname_index" ON "doctor"("surname");

ALTER TABLE "doctor" ADD CONSTRAINT "doctor_id_unique" UNIQUE("doctor_id"); 

ALTER TABLE "doctor" ADD CONSTRAINT "doctor_doctor_id_foreign" FOREIGN KEY("doctor_id") REFERENCES "user"("id");