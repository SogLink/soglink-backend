CREATE TABLE IF NOT EXISTS "file"(
    "id" SERIAL,
    "guid" UUID NOT NULL,
    "path" TEXT NOT NULL,
    "created_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "file" ADD PRIMARY KEY("id");

ALTER TABLE "file" ADD CONSTRAINT "file_guid_unique" UNIQUE("guid");

ALTER TABLE "attachment" ADD CONSTRAINT "attachment_file_id_foreign" FOREIGN KEY("file_id") REFERENCES "file"("id");