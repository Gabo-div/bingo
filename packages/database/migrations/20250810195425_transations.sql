-- +goose Up
-- +goose StatementBegin
create table "bcv_price" (
  "id" bigserial primary key, 
  "price" decimal not null, 
  "data" timestamp not null
);

create table "transaction_rules" (
  "id" bigserial primary key, 
  "min_deposit_coin" decimal not null,
  "min_withdraw_coin" decimal not null,
  "createdAt" timestamp not null 
);

create table "payment_method_type" (
  "id" bigserial primary key, 
  "name" text not null,
  "shape" jsonb not null
);

create table "payment_method" (
  "id" bigserial primary key, 
  "userId" bigserial not null references "user" ("id"), 
  "methodTypeId" bigserial not null references "payment_method_type" ("id"), 
  "data" jsonb not null 
);

create table "transaction" (
  "id" bigserial primary key, 
  "userId" bigserial not null references "user" ("id"), 
  "paymentMethodId" bigserial not null references "payment_method" ("id"), 
  "amount" decimal not null, 
  "createdAt" timestamp not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table "transaction";
drop table "payment_method";
drop table "payment_method_type";
drop table "transaction_rules";
drop table "bcv_price";
-- +goose StatementEnd
