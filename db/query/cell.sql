-- name: CreateCell :one
INSERT INTO cells (
    cell_name,
    cell_name_old,
    cell_id_givin,
    cell_id_givin_old,
    sector_name,
    uplinkuarfcn,
    downlinkuarfcn,
    dlprscramblecode,
    azimuth,
    height,
    etilt,
    mtilt,
    antennatype,
    antennamodel,
    ecgi,
    site_id,
    carrier_id,
    serviceareatype_id
) VALUES (
 $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18
)
RETURNING *;

-- name: GetCell0 :one
SELECT * FROM cells
WHERE id = $1 LIMIT 1;

-- name: GetCell1 :one
SELECT * FROM cells
WHERE cell_name = $1 LIMIT 1;

-- name: ListCells :many
SELECT * FROM cells
WHERE site_id = $3 OR carrier_id = $4 OR serviceareatype_id = $5
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateCell :one
UPDATE cells
SET cell_name = $2,
cell_name_old = $3,
cell_id_givin = $4,
cell_id_givin_old = $5,
sector_name = $6,
uplinkuarfcn = $7,
downlinkuarfcn = $8,
dlprscramblecode = $9,
azimuth = $10,
height = $11,
etilt = $12,
mtilt = $13,
antennatype = $14,
antennamodel = $15,
ecgi = $16,
site_id = $17,
carrier_id = $18,
serviceareatype_id = $19
WHERE id = $1
RETURNING *;

-- name: DeleteCell :exec
DELETE FROM cells
WHERE id = $1;
