CREATE TABLE IF NOT EXISTS "emr"(
    "id" SERIAL,
    "doctor_id" BIGINT NOT NULL,
    "patient_id" BIGINT NOT NULL,
    "diagnoses_text" TEXT NOT NULL,
    "prescriptions_text" TEXT NOT NULL,
    "created_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "emr" ADD PRIMARY KEY("id");

ALTER TABLE  "appointment" ADD CONSTRAINT "appointment_emr_id_foreign" FOREIGN KEY("emr_id") REFERENCES "emr"("id");

ALTER TABLE "emr" ADD CONSTRAINT "emr_doctor_id_foreign" FOREIGN KEY("doctor_id") REFERENCES "user"("id");

ALTER TABLE "emr" ADD CONSTRAINT "emr_patient_id_foreign" FOREIGN KEY("patient_id") REFERENCES "user"("id");