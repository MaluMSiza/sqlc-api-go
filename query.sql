-- name: GetUser :one
SELECT * FROM "user" WHERE id = $1;

-- name: InsertUser :one
INSERT INTO "user" (name, birth_date, cpf)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateUser :one
UPDATE "user"
SET name = $1, birth_date = $2, cpf = $3, updated_at = NOW()
WHERE id = $4
RETURNING *;

-- name: DeleteUser :one
DELETE FROM "user" WHERE id = $1
RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM "user";

-- name: GetUserAge :one
SELECT DATE_PART('year', AGE(birth_date)) AS age
FROM "user"
WHERE id = $1;
