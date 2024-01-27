CREATE TABLE IF NOT EXISTS "clinic"(
    "id" SERIAL,
    "guid" UUID NOT NULL,
    "location_id" BIGINT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "latitude" DOUBLE PRECISION NOT NULL,
    "longitude" DOUBLE PRECISION NOT NULL, 
    "created_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE
    "clinic"
ADD
    PRIMARY KEY("id");

ALTER TABLE
    "clinic"
ADD
    CONSTRAINT "clinic_guid_unique" UNIQUE("guid");

ALTER TABLE
    "doctor"
ADD
    CONSTRAINT "doctor_clinic_id_foreign" FOREIGN KEY("clinic_id") REFERENCES "clinic"("id");