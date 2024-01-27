CREATE TABLE IF NOT EXISTS "appointment"(
    "id" SERIAL,
    "guid" UUID NOT NULL,
    "doctor_id" BIGINT NOT NULL,
    "patient_id" BIGINT NOT NULL,
    "appointment_at" TIMESTAMP(0) WITH TIME zone NOT NULL,
    "appointment_reason" TEXT NOT NULL,
    "status" VARCHAR(255) NOT NULL,
    "emr_id" BIGINT NOT NULL,
    "created_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "appointment" ADD PRIMARY KEY("id");

ALTER TABLE "appointment" ADD CONSTRAINT "appointment_guid_unique" UNIQUE("guid");

ALTER TABLE "appointment" ADD CONSTRAINT "appointment_patient_id_foreign" FOREIGN KEY("patient_id") REFERENCES "user"("id");

ALTER TABLE "appointment" ADD CONSTRAINT "appointment_doctor_id_foreign" FOREIGN KEY("doctor_id") REFERENCES "user"("id");