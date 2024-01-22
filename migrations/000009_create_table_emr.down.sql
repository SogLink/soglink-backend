ALTER TABLE  "appointment" DROP CONSTRAINT "appointment_emr_id_foreign";

ALTER TABLE "emr" DROP CONSTRAINT "emr_doctor_id_foreign";

ALTER TABLE "emr" DROP CONSTRAINT "emr_patient_id_foreign";

DROP TABLE IF EXISTS "emr";