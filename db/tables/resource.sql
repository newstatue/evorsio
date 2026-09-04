CREATE TABLE resource (
                          id TEXT PRIMARY KEY,
                          type TEXT NOT NULL,
                          name TEXT NOT NULL,
                          created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_resource_type ON resource(type);


CREATE TABLE operation (
                           id TEXT PRIMARY KEY,
                           operation_type TEXT NOT NULL,
                           resource_id TEXT NOT NULL,

                           old_parent_id TEXT,
                           new_parent_id TEXT,

                           old_name TEXT,
                           new_name TEXT,

                           created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE operation_state (
                                 id INTEGER PRIMARY KEY CHECK (id = 1),
                                 cursor INTEGER NOT NULL DEFAULT 0
);