-- name: SelectResourceById :one
select * from resource where id = ?;

-- name: SelectResourcesByName :many
select * from resource where name like '%'||sqlc.arg(name) ||'%';

-- name: InsertResource :exec
insert into resource (id, type, name, created_at, updated_at) values (?,?,?,?,?);

-- name: UpdateResourceById :execresult
update resource set name = coalesce(sqlc.narg(name),name), updated_at = ? where id = ?;

-- name: DeleteResourceById :execresult
delete from resource where id = ?;

-- name: DeleteResourcesByIds :execresult
delete from resource where id in (sqlc.slice(ids));