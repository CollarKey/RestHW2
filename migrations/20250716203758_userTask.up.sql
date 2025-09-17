ALTER TABLE tasks ADD COLUMN user_id INTEGER REFERENCES users(id)
ON DELETE CASCADE;

ALTER TABLE tasks
ADD CONSTRAINT fk_tasks_users
FOREIGN KEY (user_id)
REFERENCES users(id)
ON DELETE CASCADE;

CREATE INDEX idx_tasks_user_id ON tasks(user_id);
