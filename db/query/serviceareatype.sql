-- name: CreateServiceareatype :one
INSERT INTO serviceareatypes (
    serviceareatype_name,
    serviceareatype_desc
) VALUES (
 $1, $2
)
RETURNING *;

-- name: GetServiceareatype0 :one
SELECT * FROM serviceareatypes
WHERE id = $1 LIMIT 1;

-- name: GetServiceareatype1 :one
SELECT * FROM serviceareatypes
WHERE serviceareatype_name = $1 LIMIT 1;

-- name: ListServiceareatypes :many
SELECT * FROM serviceareatypes
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateServiceareatype :one
UPDATE serviceareatypes
SET serviceareatype_name = $2,
serviceareatype_desc = $3
WHERE id = $1
RETURNING *;

-- name: DeleteServiceareatype :exec
DELETE FROM serviceareatypes
WHERE id = $1;
