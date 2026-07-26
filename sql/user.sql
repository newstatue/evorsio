CREATE TABLE users (
                       id UUID PRIMARY KEY,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       name VARCHAR(100),
                       status SMALLINT NOT NULL DEFAULT 0,
                       created_at TIMESTAMPTZ NOT NULL,
                       updated_at TIMESTAMPTZ NOT NULL
);

COMMENT ON TABLE users IS '用户表';

COMMENT ON COLUMN users.id IS '用户唯一标识（UUID）';
COMMENT ON COLUMN users.email IS '用户邮箱';
COMMENT ON COLUMN users.name IS '用户昵称';
COMMENT ON COLUMN users.status IS '用户状态：0-正常，1-禁用';
COMMENT ON COLUMN users.created_at IS '创建时间';
COMMENT ON COLUMN users.updated_at IS '更新时间';