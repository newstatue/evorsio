CREATE TABLE objects (
                         id UUID PRIMARY KEY,

                         bucket VARCHAR(255) NOT NULL,
                         object_key TEXT NOT NULL,

                         sha256 VARCHAR(64),
                         size BIGINT NOT NULL DEFAULT 0,
                         content_type VARCHAR(255) NOT NULL DEFAULT 'application/octet-stream',
                         etag VARCHAR(255),

                         storage_class VARCHAR(64) NOT NULL DEFAULT 'STANDARD',
                         ref_count BIGINT NOT NULL DEFAULT 0,
                         status VARCHAR(20) NOT NULL DEFAULT 'pending',

                         created_at TIMESTAMPTZ NOT NULL,
                         updated_at TIMESTAMPTZ NOT NULL,
                         deleted_at TIMESTAMPTZ,

                         UNIQUE (bucket, object_key)
);

COMMENT ON TABLE objects IS '物理存储对象表';

COMMENT ON COLUMN objects.id IS '对象唯一标识';
COMMENT ON COLUMN objects.bucket IS '对象存储桶名称';
COMMENT ON COLUMN objects.object_key IS '对象在存储桶中的Key';
COMMENT ON COLUMN objects.sha256 IS '文件内容SHA-256摘要';
COMMENT ON COLUMN objects.size IS '文件大小，单位字节';
COMMENT ON COLUMN objects.content_type IS '文件MIME类型';
COMMENT ON COLUMN objects.etag IS '对象存储返回的ETag';
COMMENT ON COLUMN objects.storage_class IS '对象存储类型';
COMMENT ON COLUMN objects.ref_count IS '对象引用次数';
COMMENT ON COLUMN objects.status IS '状态：pending、active、deleting、deleted';
COMMENT ON COLUMN objects.created_at IS '创建时间';
COMMENT ON COLUMN objects.updated_at IS '更新时间';
COMMENT ON COLUMN objects.deleted_at IS '删除时间';

CREATE INDEX idx_objects_sha256
    ON objects (sha256);

CREATE INDEX idx_objects_status
    ON objects (status);


