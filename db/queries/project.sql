-- name: AddProject :exec
INSERT OR IGNORE INTO projects (name) VALUES (?);

-- name: GetProject :one
SELECT name FROM projects WHERE name = ?;
