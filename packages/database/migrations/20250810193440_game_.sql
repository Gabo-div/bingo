-- +goose Up
-- +goose StatementBegin
create table "game_type" (
  "id" bigserial primary key, 
  "name" text not null
);

create table "game" (
  "id" bigserial primary key, 
  "name" text not null, 
  "start" timestamp not null,
  "open" boolean not null,
  "cardPrice" decimal not null,
  "maxPlayers" serial not null,
  "gameTypeId" bigserial not null references "game_type" ("id"), 
  "createdAt" timestamp not null, 
  "updatedAt" timestamp not null
);

create table "game_ball" (
  "id" bigserial primary key, 
  "gameId" bigserial not null references "game" ("id"), 
  "number" smallserial not null, 
  "createdAt" timestamp not null
);

create table "game_card" (
  "id" bigserial primary key, 
  "userId" bigserial not null references "user" ("id"), 
  "gameId" bigserial not null references "game" ("id"), 
  "winner" boolean not null, 
  "seed" bigserial not null, 
  "createdAt" timestamp not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table "game_card";
drop table "game_ball";
drop table "game";
drop table "game_type";
-- +goose StatementEnd
