-- +goose Up

INSERT INTO projects("name")
VALUES ("Portfolio");

INSERT INTO tasks("name", "priority", "project")
VALUES  
  ("Create Landing Page", "H", "Portfolio"),
  ("Setup DB", "M", ""),
  ("Update Readme", "L", "");
-- +goose Down
DELETE FROM tasks;
DELETE FROM projects;
