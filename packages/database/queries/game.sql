-- name: GetGameByID :one
SELECT * FROM "game" WHERE "id" = $1 LIMIT 1;

-- name: ListGames :many
SELECT * FROM "game" ORDER BY "createdAt" DESC;

-- name: GetGameBalls :many
SELECT * FROM "game_ball" WHERE "gameId" = $1 ORDER BY "createdAt" ASC;

-- name: GetGameCards :many
SELECT * FROM "game_card" WHERE "gameId" = $1;

-- name: GetGameCardByUser :one
SELECT * FROM "game_card" WHERE "gameId" = $1 AND "userId" = $2;

-- name: GetGameWinner :one
SELECT * FROM "game_card" WHERE "gameId" = $1 AND "winner" = true;
