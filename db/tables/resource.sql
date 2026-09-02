CREATE TABLE resource (
                          id TEXT PRIMARY KEY,
                          type TEXT NOT NULL,
                          name TEXT NOT NULL,
                          created_at TEXT NOT NULL,
                          updated_at TEXT NOT NULL
);

CREATE INDEX idx_resource_type ON resource(type);