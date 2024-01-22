CREATE TABLE IF NOT EXISTS "attachment"(
    "emr_id" BIGINT NOT NULL,
    "file_id" BIGINT NOT NULL,
    "created_at" TIMESTAMP(0) WITH TIME zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "attachment" ADD PRIMARY KEY("emr_id", "file_id");

ALTER TABLE "attachment" ADD CONSTRAINT "attachment_emr_id_foreign" FOREIGN KEY("emr_id") REFERENCES "emr"("id");
