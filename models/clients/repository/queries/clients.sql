
-- name: CreateClient :one
INSERT INTO clients(
  fanfit_user_id,
  temp_field
) VALUES(
  $1,
  $2
)
RETURNING *;


-- name: DeleteUser :one
DELETE FROM users
WHERE email = $1
RETURNING *;


-- name: GetClientByEmail :one
SELECT * FROM users INNER JOIN clients
ON users.id = clients.fanfit_user_id
WHERE email = $1;

-- name: GetClientByID :one
SELECT * FROM users INNER JOIN clients
ON users.id = clients.fanfit_user_id
WHERE fanfit_user_id = $1;
