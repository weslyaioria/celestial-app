-- name: CreateAccount :one
INSERT INTO productCategory ( 
  name
) VALUES (
  $1
) RETURNING *;
