-- +goose Up
CREATE TABLE projects(
  name VARCHAR(100) UNIQUE,
  rate REAL DEFAULT 0.00 
);

CREATE TABLE tasks(
  id INTEGER PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  project VARCHAR(100), 
  client VARCHAR(255),
  priority VARCHAR(3),
  notes TEXT,
  due_date DATETIME,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (project)
      REFERENCES projects(name)
);


CREATE TABLE sessions(
  id INTEGER PRIMARY KEY,
  task_id INTEGER NOT NULL,
  start DATETIME DEFAULT CURRENT_TIMESTAMP,
  stop DATETIME,
  FOREIGN KEY (task_id)
       REFERENCES tasks(id)
);
-- +goose Down
DROP TABLE sessions;
DROP TABLE tasks;
DROP TABLE projects;
