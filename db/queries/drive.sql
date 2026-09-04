-- name: SelectFileById :one
select
    sqlc.embed(r),
    sqlc.embed(f)
from file f
         join resource r
                    on r.id = f.resource_id
where r.id = ? and r.type = 'file';

-- name: InsertFile :exec
insert into file (resource_id, size, mime_type) values (?,?,?);

-- name: SelectFolderById :one
select
    sqlc.embed(r)
from resource r where r.id = ? and r.type = 'folder';

-- name: SelectSymlinkById :one
select
    sqlc.embed(r),
    sqlc.embed(s)
from symlink s
         join resource r
                    on r.id = s.resource_id
where r.id = ? and r.type = 'symlink';

-- name: InsertSymlink :exec
insert into symlink (resource_id, target_id) values (?,?);

-- name: InsertEntry :exec
insert into entry (parent_id, child_id) values (?,?);

-- name: SelectParentByChildId :one
select sqlc.embed(r) from entry e
join resource r on r.id = e.parent_id
where e.child_id = ?;

-- name: SelectChildrenByParentId :many
select sqlc.embed(r) from entry e
join resource r on r.id = e.child_id
where e.parent_id =?
order by r.updated_at desc;

-- name: UpdateEntryParentByChildId :exec
update entry set parent_id = ? where child_id = ?;

-- name: ExistsFolderById :one
select exists(select 1 from resource where id = ? and type = 'folder')