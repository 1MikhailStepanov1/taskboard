-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS projects(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    owner UUID NOT NULL,
    is_archived BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS boards(
    id VARCHAR(50) PRIMARY KEY,
    project_id INT NOT NULL,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL PRIMARY KEY,
    board_id VARCHAR(50) NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    status VARCHAR(50) NOT NULL,
    priority VARCHAR(20) NOT NULL,
    created_by UUID NOT NULL,
    assigned UUID,
    deadline TIMESTAMP
);

CREATE TABLE IF NOT EXISTS comments(
    id UUID PRIMARY KEY,
    task_id VARCHAR(100) NOT NULL,
    user_id UUID NOT NULL,
    text TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE commets;
DROP TABLE tasks;
DROP TABLE boards;
DROP TABLE projects;
-- +goose StatementEnd
