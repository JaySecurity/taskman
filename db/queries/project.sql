-- name: AddProject :one
INSERT OR IGNORE INTO projects (name) VALUES (?) RETURNING *;

-- name: GetProject :one
SELECT * FROM projects WHERE name = ?;

-- name: ModifyProject :one
UPDATE projects
SET rate = COALESCE(sqlc.arg(rate), rate),
    name = sqlc.arg(new_name)
WHERE name = sqlc.arg(name)
RETURNING *;

-- name: DeleteProject :exec
DELETE FROM projects WHERE name = ?;
