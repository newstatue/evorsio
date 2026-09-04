CREATE TABLE file (
                      resource_id TEXT PRIMARY KEY,
                      size INTEGER NOT NULL DEFAULT 0,
                      mime_type TEXT NOT NULL DEFAULT '',
                      FOREIGN KEY (resource_id) REFERENCES resource(id) ON DELETE CASCADE
);

CREATE TABLE symlink (
                         resource_id TEXT PRIMARY KEY,
                         target_id TEXT NOT NULL,
                         FOREIGN KEY (resource_id) REFERENCES resource(id) ON DELETE CASCADE
);

CREATE TABLE entry (
                       parent_id TEXT NOT NULL,
                       child_id TEXT NOT NULL,

                       FOREIGN KEY (parent_id) REFERENCES resource(id) ON DELETE CASCADE,
                       FOREIGN KEY (child_id) REFERENCES resource(id) ON DELETE CASCADE
);