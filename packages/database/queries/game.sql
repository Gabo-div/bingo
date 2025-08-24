-- name: GetGame :one
SELECT * FROM game WHERE "id" = $1;

-- name: CreateGame :one
INSERT INTO game (
  "name", 
  "start", 
  "open", 
  "cardPrice", 
  "maxPlayers", 
  "gameTypeId", 
  "createdAt", 
  "updatedAt"
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: EditGame :one
UPDATE game SET 
  "name" = coalesce(sqlc.narg('name'), "name"), 
  "start" = coalesce(sqlc.narg('start'), "start"), 
  "open" = coalesce(sqlc.narg('open'), "open"),
  "cardPrice" = coalesce(sqlc.narg('cardPrice'), "cardPrice"),
  "maxPlayers" = coalesce(sqlc.narg('maxPlayers'), "maxPlayers"),
  "gameTypeId" = coalesce(sqlc.narg('gameTypeId'), "gameTypeId"),
  "createdAt" = coalesce(sqlc.narg('createdAt'), "createdAt"), 
  "updatedAt" = coalesce(sqlc.narg('updatedAt'), "updatedAt")
 WHERE "id" = $1 RETURNING *;

-- name: DeleteGame :one
DELETE FROM game WHERE "id" = $1 RETURNING *;
