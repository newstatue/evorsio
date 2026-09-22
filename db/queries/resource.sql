-- name: SelectResourceById :one
select * from resource where id = ?;

-- name: SelectResourcesByPath :many
select * from resource where path like '%'||sqlc.arg(path) ||'%';

-- name: InsertResource :exec
insert into resource (id, kind,path, created_at, updated_at) values (?,?,?,?,?);

-- name: DeleteResourceById :execresult
delete from resource where id = ?;

-- name: DeleteResourcesByIds :execresult
delete from resource where id in (sqlc.slice(ids));