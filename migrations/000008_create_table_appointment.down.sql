ALTER TABLE "appointment" DROP CONSTRAINT "appointment_patient_id_foreign";

ALTER TABLE "appointment" DROP CONSTRAINT "appointment_doctor_id_foreign"; 

DROP TABLE IF EXISTS "appointment";