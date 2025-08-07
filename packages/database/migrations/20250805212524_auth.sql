-- +goose Up
-- +goose StatementBegin
create table "user" (
  "id" bigserial primary key, 
  "name" text not null, 
  "email" text not null unique, 
  "emailVerified" boolean not null, 
  "role" text,
  "banned" boolean,
  "banReason" text,
  "banExpires" timestamp,
  "image" text, "createdAt" timestamp not null, 
  "updatedAt" timestamp not null
);

create table "session" (
  "id" bigserial primary key, 
  "expiresAt" timestamp null, 
  "token" text not null unique, 
  "impersonatedBy" text,
  "createdAt" timestamp not null, 
  "updatedAt" timestamp not null, 
  "ipAddress" text, 
  "userAgent" text, 
  "userId" bigserial not null references "user" ("id")
);

create table "account" (
  "id" bigserial primary key, 
  "accountId" text not null, 
  "providerId" text not null, 
  "userId" bigserial not null references "user" ("id"), 
  "accessToken" text, 
  "refreshToken" text, 
  "idToken" text, 
  "accessTokenExpiresAt" timestamp, 
  "refreshTokenExpiresAt" timestamp, 
  "scope" text, 
  "password" text, 
  "createdAt" timestamp not null, 
  "updatedAt" timestamp not null
);

create table "verification" (
  "id" bigserial primary key, 
  "identifier" text not null, 
  "value" text not null, 
  "expiresAt" timestamp not null, 
  "createdAt" timestamp, 
  "updatedAt" timestamp
);

create table "jwks" (
  "id" bigserial primary key, 
  "publicKey" text not null, 
  "privateKey" text not null, 
  "createdAt" timestamp not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table "jwks";
drop table "verification";
drop table "account";
drop table "session";
drop table "user";
-- +goose StatementEnd
