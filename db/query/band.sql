-- name: CreateBand :one
INSERT INTO bands (
    band_name,
    band_desc,
    size,
    start_freq,
    end_freq,
    tech_id
) VALUES (
 $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetBand0 :one
SELECT * FROM bands
WHERE id = $1 LIMIT 1;

-- name: GetBand1 :one
SELECT * FROM bands
WHERE band_name = $1 LIMIT 1;

-- name: ListBands :many
SELECT * FROM bands
WHERE tech_id = $3
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateBand :one
UPDATE bands
SET band_name = $2,
band_desc = $3,
size = $4,
start_freq = $5,
end_freq = $6,
tech_id = $7
WHERE id = $1
RETURNING *;

-- name: DeleteBand :exec
DELETE FROM bands
WHERE id = $1;
