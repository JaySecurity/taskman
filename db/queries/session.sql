-- name: GetSession :one
SELECT *
FROM sessions
WHERE id = ?;

-- name: GetTaskSessions :many
SELECT *
FROM sessions
WHERE task_id = ?
ORDER BY start DESC;


-- name: Start :one
INSERT INTO sessions (task_id, start,stop )
VALUES (?, ?, ?)
RETURNING *;


-- name: Stop :one
UPDATE sessions
SET stop = ?
WHERE id = ?
RETURNING *;

-- name: ModifySession :one
UPDATE sessions
SET start = COALESCE(?, start),
  stop = COALESCE(?, stop)
WHERE id = ?
RETURNING *;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = ?;
