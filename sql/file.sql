CREATE TABLE files (
                       id UUID PRIMARY KEY,

                       owner_id UUID NOT NULL,
                       parent_id UUID,
                       object_id UUID,

                       name VARCHAR(255) NOT NULL,
                       type VARCHAR(20) NOT NULL,
                       status VARCHAR(20) NOT NULL DEFAULT 'active',

                       created_at TIMESTAMPTZ NOT NULL,
                       updated_at TIMESTAMPTZ NOT NULL,
                       deleted_at TIMESTAMPTZ
);

COMMENT ON TABLE files IS '文件和目录元数据表';

COMMENT ON COLUMN files.id IS '文件或目录唯一标识';
COMMENT ON COLUMN files.owner_id IS '所有者用户ID';
COMMENT ON COLUMN files.parent_id IS '父目录ID，根目录下为空';
COMMENT ON COLUMN files.object_id IS '关联的物理对象ID，目录为空';
COMMENT ON COLUMN files.name IS '文件或目录名称';
COMMENT ON COLUMN files.type IS '类型：file、directory';
COMMENT ON COLUMN files.status IS '状态：pending、active、deleted';
COMMENT ON COLUMN files.created_at IS '创建时间';
COMMENT ON COLUMN files.updated_at IS '更新时间';
COMMENT ON COLUMN files.deleted_at IS '删除时间';

CREATE INDEX idx_files_owner_parent
    ON files (owner_id, parent_id);

CREATE INDEX idx_files_object_id
    ON files (object_id);

CREATE INDEX idx_files_status
    ON files (status);