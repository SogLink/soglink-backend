CREATE TABLE IF NOT EXISTS "location"(
    "id" SERIAL,
    "city" VARCHAR(255) NOT NULL,
    "region" VARCHAR(255) NOT NULL,
    "latitude" DOUBLE PRECISION NOT NULL,
    "longitude" DOUBLE PRECISION NOT NULL
);
ALTER TABLE "location" ADD PRIMARY KEY("id");

ALTER TABLE "location" ADD CONSTRAINT "location_city_region_unique" UNIQUE("city", "region");

ALTER TABLE "clinic" ADD CONSTRAINT "clinic_location_id_foreign" FOREIGN KEY("location_id") REFERENCES "location"("id");