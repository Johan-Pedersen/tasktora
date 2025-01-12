CREATE TABLE  tasks(
  id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(100) NOT NULL,
  note TEXT NOT NULL,
  created DATETIME NOT NULL,
  parent_id INT NULL,
  level INT NOT NULL
);
CREATE INDEX idx_task_created ON tasks(created);

INSERT INTO tasks(title, note, created, level) VALUES(
  "fugl 1",
  "Synger",
  '2025-01-01 10:00:00',
  1
  );
INSERT INTO tasks(title, note, created, level) VALUES(
  "fugl 2",
  "flyver",
  '2025-01-01 10:00:00',
  1
  );
INSERT INTO tasks(title, note, created, parent_id, level) VALUES(
  "fugl 3",
  "Gaar",
  '2025-01-01 10:00:00',
  1,
  2
  );
INSERT INTO tasks(title, note, created, parent_id, level) VALUES(
  "fugl 4",
  "danser",
  '2025-01-01 10:00:00',
  3,
  3
  );
INSERT INTO tasks(title, note, created, parent_id, level) VALUES(
  "fugl 5",
  "sover",
  '2025-01-01 10:00:00',
  1,
  2
  );
