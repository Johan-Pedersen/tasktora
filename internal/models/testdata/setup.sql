CREATE TABLE  tasks(
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(100) NOT NULL,
  note TEXT NOT NULL,
  created DATETIME NOT NULL,
  parent_id int NULL
);
CREATE INDEX idx_task_created ON tasks(created);

INSERT INTO tasks(title, note, created) VALUES(
  "fugl 1",
  "Synger",
  '2025-01-01 10:00:00'
  );
INSERT INTO tasks(title, note, created) VALUES(
  "fugl 2",
  "flyver",
  '2025-01-01 10:00:00'
  );
INSERT INTO tasks(title, note, created, parent_id) VALUES(
  "fugl 3",
  "Gaar",
  '2025-01-01 10:00:00',
  1
  );
INSERT INTO tasks(title, note, created, parent_id) VALUES(
  "fugl 4",
  "danser",
  '2025-01-01 10:00:00',
  3
  );
INSERT INTO tasks(title, note, created, parent_id) VALUES(
  "fugl 5",
  "sover",
  '2025-01-01 10:00:00',
  1
  );
