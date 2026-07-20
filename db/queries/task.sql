-- name: GetTask :one
SELECT
  *
FROM
  tasks
WHERE
  id = ?
LIMIT 1;


-- name: GetAllTasks :many
SELECT
  *
FROM
  tasks
ORDER BY
  due_date;


-- name: CreateTask :one
INSERT INTO tasks ("name", "project", "client", "priority", "due_date" )
VALUES(?,?,?,?,?)
RETURNING *;


-- name: ModifyTask :one
UPDATE tasks
SET
  "name" = COALESCE(?, "name"), 
  "project" = COALESCE(?, "project"),
  "client" = COALESCE(?, "client"),
  "priority" = COALESCE(?, "priority"),
  "due_date" = COALESCE(?, "due_date"),
  "current_session" = COALESCE(?, "current_session")
WHERE
  id = ?
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE
  id = ?;
