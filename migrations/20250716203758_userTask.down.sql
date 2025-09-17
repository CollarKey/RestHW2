ALTER TABLE tasks DROP CONSTRAINT fk_tasks_users;

DROP INDEX idx_tasks_user_id;

ALTER TABLE tasks DROP COLUMN user_id;
