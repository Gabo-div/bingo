-- name: GetUserByID :one
SELECT * FROM "user" WHERE "id" = $1 LIMIT 1;

-- name: AddPaymentMethodWithUserID :exec
INSERT INTO "payment_method" ("userId", "methodTypeId", "data") VALUES ($1, $2, $3);

-- name: GetPaymentMethodsByUserID :many
SELECT * FROM "payment_method" WHERE "userId" = $1;

-- name: GetTransactionHistoryByUserID :many
SELECT * FROM "transaction" WHERE "userId" = $1;
