CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "address" varchar,
  "phone" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp
);

CREATE TABLE "categories" (
  "id" uuid,
  "name" varchar UNIQUE NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp
);

CREATE TABLE "products" (
  "id" uuid,
  "name" varchar NOT NULL,
  "description" text NOT NULL,
  "price" decimal NOT NULL,
  "stock" integer NOT NULL,
  "category_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp
);

CREATE TABLE "orders" (
  "id" uuid,
  "user_id" uuid NOT NULL,
  "total_amount" decimal NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp
);

CREATE TABLE "order_items" (
  "id" uuid PRIMARY KEY,
  "order_id" uuid NOT NULL,
  "product_id" uuid NOT NULL,
  "quantity" integer NOT NULL,
  "price" decimal NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "charts" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "chart_items" (
  "id" uuid PRIMARY KEY,
  "cart_id" uuid NOT NULL,
  "product_id" uuid NOT NULL,
  "quantity" integer NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
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
