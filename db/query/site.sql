-- name: CreateSite :one
INSERT INTO sites (
    site_name,
    site_name_old,
    site_id_givin,
    site_id_givin_old,
    lac,
    rac,
    rnc,
    site_on_air_date,
    property_id,
    sitetype_id,
    vendor_id
) VALUES (
 $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- name: GetSite0 :one
SELECT * FROM sites
WHERE id = $1 LIMIT 1;

-- name: GetSite1 :one
SELECT * FROM sites
WHERE site_name = $1 LIMIT 1;

-- name: ListSites :many
SELECT * FROM sites
WHERE property_id = $3 OR sitetype_id = $4 OR vendor_id = $5
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateSite :one
UPDATE sites
SET site_name = $2,
site_name_old = $3,
site_id_givin = $4,
site_id_givin_old = $5,
lac = $6,
rac = $7,
rnc = $8,
site_on_air_date = $9,
property_id = $10,
sitetype_id = $11,
vendor_id = $12
WHERE id = $1
RETURNING *;

-- name: DeleteSite :exec
DELETE FROM sites
WHERE id = $1;
