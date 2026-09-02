-- name: SelectFileById :one
select
    sqlc.embed(r),
    sqlc.embed(f)
from file f
         join resource r
                    on r.id = f.resource_id
where r.id = ? and r.type = 'file';


-- name: SelectFolderById :one
select
    sqlc.embed(r)
from resource r where id = ? and type = 'folder';

-- name: SelectSymlinkById :one
select
    sqlc.embed(r),
    sqlc.embed(s)
from symlink s
         join resource r
                    on r.id = s.resource_id
where r.id = ? and r.type = 'symlink';