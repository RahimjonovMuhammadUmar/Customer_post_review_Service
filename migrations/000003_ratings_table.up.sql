CREATE TABLE "ratings"(
    "id" SERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "rating" INTEGER NOT NULL,
    "description" TEXT NOT NULL,
    "post_id" INTEGER NOT NULL,
    "customer_id" INTEGER NOT NULL,
    "deleted_at" TIMESTAMP(0) WITH
        TIME zone NULL
);