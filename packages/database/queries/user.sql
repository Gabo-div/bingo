-- name: GetUserByID :one
SELECT * FROM "user" WHERE "id" = $1 LIMIT 1;
