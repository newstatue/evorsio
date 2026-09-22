-- +goose Up
SELECT 'up SQL query';
CREATE TABLE resource (
                          id TEXT PRIMARY KEY,
                          kind TEXT NOT NULL,
                            path TEXT NOT NULL UNIQUE ,
                          created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE operation (
                           id TEXT PRIMARY KEY,
                           operation_type TEXT NOT NULL,
                           resource_id TEXT NOT NULL,

                           old_path TEXT,
                           new_path TEXT,

                           created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        FOREIGN KEY (resource_id) REFERENCES resource(id) ON DELETE CASCADE
);

-- +goose Down
SELECT 'down SQL query';
DROP TABLE IF EXISTS resource;
DROP TABLE IF EXISTS operation;
