CREATE INDEX idx_tasks_author_user_id_created_at
    ON todoapp.tasks(author_user_id, created_at);
