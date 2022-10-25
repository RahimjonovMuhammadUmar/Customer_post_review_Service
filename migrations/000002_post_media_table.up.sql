CREATE TABLE "posts"(
    "id" SERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "created_at" TIMESTAMP(0) WITH TIME zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP(0) WITH TIME zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "deleted_at" TIMESTAMP(0) WITH TIME zone NULL,
    "customer_id" INTEGER NOT NULL
);

CREATE TABLE "medias"(
    "id" SERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "link" TEXT NOT NULL,
    "type" TEXT NOT NULL,
    "post_id" INTEGER NOT NULL,
    "deleted_at" TIMESTAMP(0) WITH
        TIME zone NULL
);
ALTER TABLE
    "medias" ADD CONSTRAINT "medias_post_id_foreign" FOREIGN KEY("post_id") REFERENCES "posts"("id");
