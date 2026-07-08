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
