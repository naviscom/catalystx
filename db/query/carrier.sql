-- name: CreateCarrier :one
INSERT INTO carriers (
    carrier_name,
    carrier_desc,
    size,
    start_freq,
    end_freq,
    band_id
) VALUES (
 $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetCarrier0 :one
SELECT * FROM carriers
WHERE id = $1 LIMIT 1;

-- name: GetCarrier1 :one
SELECT * FROM carriers
WHERE carrier_name = $1 LIMIT 1;

-- name: ListCarriers :many
SELECT * FROM carriers
WHERE band_id = $3
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateCarrier :one
UPDATE carriers
SET carrier_name = $2,
carrier_desc = $3,
size = $4,
start_freq = $5,
end_freq = $6,
band_id = $7
WHERE id = $1
RETURNING *;

-- name: DeleteCarrier :exec
DELETE FROM carriers
WHERE id = $1;
