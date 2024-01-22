CREATE TABLE IF NOT EXISTS "patient"(
    "patient_id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "surname" VARCHAR(255) NOT NULL,
    "gender" VARCHAR(255) NOT NULL,
    "birthday" DATE NOT NULL,
    "pinfl" BIGINT NOT NULL
);

CREATE INDEX "patient_name_index" ON "patient"("name");
CREATE INDEX "patient_surname_index" ON "patient"("surname");
CREATE INDEX "patient_pinfl_index" ON "patient"("pinfl");
ALTER TABLE "patient" ADD CONSTRAINT "patient_id_unique" UNIQUE("patient_id");  


ALTER TABLE "patient" ADD CONSTRAINT "patient_user_id_foreign" FOREIGN KEY("patient_id") REFERENCES "user"("id");