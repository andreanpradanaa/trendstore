CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "address" varchar,
  "phone" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z')
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z')
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" text NOT NULL,
  "price" decimal NOT NULL,
  "stock" integer NOT NULL,
  "category_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z')
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "total_amount" decimal NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z')
);

CREATE TABLE "order_items" (
  "id" bigserial PRIMARY KEY,
  "order_id" bigserial NOT NULL,
  "product_id" bigserial NOT NULL,
  "quantity" integer NOT NULL,
  "price" decimal NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z')
);

CREATE TABLE "charts" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z')
);

CREATE TABLE "chart_items" (
  "id" bigserial PRIMARY KEY,
  "cart_id" bigserial NOT NULL,
  "product_id" bigserial NOT NULL,
  "quantity" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z')
);

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "categories" ("name");

CREATE INDEX ON "products" ("name");

CREATE INDEX ON "products" ("price");

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "charts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "chart_items" ADD FOREIGN KEY ("cart_id") REFERENCES "charts" ("id");

ALTER TABLE "chart_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
