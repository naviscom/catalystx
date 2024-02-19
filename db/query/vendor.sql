-- name: CreateVendor :one
INSERT INTO vendors (
    vendor_name,
    vendor_desc
) VALUES (
 $1, $2
)
RETURNING *;

-- name: GetVendor0 :one
SELECT * FROM vendors
WHERE id = $1 LIMIT 1;

-- name: GetVendor1 :one
SELECT * FROM vendors
WHERE vendor_name = $1 LIMIT 1;

-- name: ListVendors :many
SELECT * FROM vendors
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateVendor :one
UPDATE vendors
SET vendor_name = $2,
vendor_desc = $3
WHERE id = $1
RETURNING *;

-- name: DeleteVendor :exec
DELETE FROM vendors
WHERE id = $1;
