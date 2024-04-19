-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2024-04-19T09:51:14.997Z

CREATE TYPE "item_type" AS ENUM (
  'consumable',
  'armor',
  'rangedWeapon',
  'meleeWeapon',
  'consumableWeapon',
  'quest',
  'resource'
);

CREATE TABLE "item" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar(255),
  "text" text,
  "variant" varchar(255),
  "buy_value" int,
  "sell_value" int,
  "weight" int,
  "durability" int,
  "max_stack" int,
  "height" int,
  "width" int,
  "rawshape" text,
  "created_at" timestamp DEFAULT (now()),
  "type" item_type
);

CREATE TABLE "user" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "username" varchar(255),
  "salt" varchar(255),
  "password_hash" varchar(255),
  "email" varchar(255),
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "inventory" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "invname" varchar(255),
  "user_id" int,
  "width" int,
  "height" int,
  "max_weight" int,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "inventory_item" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "inventory_id" int,
  "item_id" int,
  "quantity" int,
  "position_x" int,
  "position_y" int,
  "rotation" int,
  "durabilityLeft" int,
  "created_at" timestamp DEFAULT (now())
);

ALTER TABLE "inventory" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "inventory_item" ADD FOREIGN KEY ("inventory_id") REFERENCES "inventory" ("id");

ALTER TABLE "inventory_item" ADD FOREIGN KEY ("item_id") REFERENCES "item" ("id");
