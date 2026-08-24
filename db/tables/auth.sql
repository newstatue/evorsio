CREATE TABLE user (
                      id TEXT PRIMARY KEY,
                      name TEXT NOT NULL,
                      email TEXT NOT NULL UNIQUE,
                      email_verified INTEGER NOT NULL DEFAULT 0,
                      image TEXT,
                      created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                      updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE session (
                         id TEXT PRIMARY KEY,
                         expires_at DATETIME NOT NULL,
                         token TEXT NOT NULL UNIQUE,
                         created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         ip_address TEXT,
                         user_agent TEXT,
                         user_id TEXT NOT NULL,
                         FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
);

CREATE TABLE account (
                         id TEXT PRIMARY KEY,
                         issuer TEXT NOT NULL,
                         account_id TEXT NOT NULL,
                         provider_id TEXT NOT NULL,
                         user_id TEXT NOT NULL,
                         access_token TEXT,
                         refresh_token TEXT,
                         id_token TEXT,
                         access_token_expires_at DATETIME,
                         refresh_token_expires_at DATETIME,
                         scope TEXT,
                         password TEXT,
                         created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
);

CREATE TABLE verification (
                              id TEXT PRIMARY KEY,
                              identifier TEXT NOT NULL,
                              value TEXT NOT NULL,
                              expires_at DATETIME NOT NULL,
                              created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_session_user_id
    ON session(user_id);

CREATE INDEX idx_account_user_id
    ON account(user_id);

CREATE UNIQUE INDEX idx_account_issuer_account_id
    ON account(issuer, account_id);

CREATE INDEX idx_verification_identifier
    ON verification(identifier);