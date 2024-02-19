-- name: CreateClutter :one
INSERT INTO clutters (
    clutter_name,
    clutter_desc
) VALUES (
 $1, $2
)
RETURNING *;

-- name: GetClutter0 :one
SELECT * FROM clutters
WHERE id = $1 LIMIT 1;

-- name: GetClutter1 :one
SELECT * FROM clutters
WHERE clutter_name = $1 LIMIT 1;

-- name: ListClutters :many
SELECT * FROM clutters
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateClutter :one
UPDATE clutters
SET clutter_name = $2,
clutter_desc = $3
WHERE id = $1
RETURNING *;

-- name: DeleteClutter :exec
DELETE FROM clutters
WHERE id = $1;
