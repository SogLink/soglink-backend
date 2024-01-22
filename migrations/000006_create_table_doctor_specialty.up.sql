CREATE TABLE IF NOT EXISTS "doctor_specialty"(
    "doctor_id" BIGINT NOT NULL,
    "specialty_id" BIGINT NOT NULL,
    "created_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "doctor_specialty" ADD PRIMARY KEY("doctor_id", "specialty_id");

ALTER TABLE "doctor_specialty" ADD CONSTRAINT "doctor_specialty_doctor_id_foreign" FOREIGN KEY("doctor_id") REFERENCES "doctor"("doctor_id");