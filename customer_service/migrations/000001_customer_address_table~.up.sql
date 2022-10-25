CREATE TABLE "customers"(
    "id" SERIAL PRIMARY KEY,
    "first_name" TEXT NOT NULL,
    "last_name" TEXT NOT NULL,
    "bio" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "created_at" TIMESTAMP(0) WITH TIME zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP(0) WITH TIME zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "deleted_at" TIMESTAMP(0) WITH TIME zone  NULL,
    "phone_number" TEXT NOT NULL
);

CREATE TABLE "addressses"(
    "id" SERIAL PRIMARY KEY,
    "house_number" INTEGER NOT NULL,
    "street" TEXT NOT NULL,
    "customer_id" INTEGER NOT NULL
);
ALTER TABLE
    "addressses" ADD CONSTRAINT "addressses_customer_id_foreign" FOREIGN KEY("customer_id") REFERENCES "customers"("id");
