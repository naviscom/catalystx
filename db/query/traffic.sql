-- name: CreateTraffic :one
INSERT INTO traffic (
    traffic_date,
    avgdailydldatamb,
    avgdailyuldatamb,
    avgdailytotdatamb,
    avgdailytotvoicemin,
    avgdailytotvideomin,
    qci1_data,
    qci6_data,
    qci8_data,
    qci_other_data,
    avgdailytotvoicemin4g,
    avgdailytotvoicemintotal,
    userdlthroughput,
    dlpacketlossrate,
    overallpsdropcallrate,
    bhdldatamb,
    bhupdatamb,
    bhtotdatamb,
    bhtotvoicemin,
    bhtotvideomin,
    bhcsusers,
    bhhsupausers,
    bhhsdpausers,
    bhr99uldl,
    powercapacity,
    powerutilization,
    codecapacity,
    codeutilization,
    ceulcapacity,
    ceulutilization,
    cedlcapacity,
    cedlutilization,
    iubcapacity,
    iubutlization,
    bhrrcusers,
    cell_id
) VALUES (
 $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36
)
RETURNING *;

-- name: GetTraffic0 :one
SELECT * FROM traffic
WHERE id = $1 LIMIT 1;

-- name: ListTraffic :many
SELECT * FROM traffic
WHERE cell_id = $3
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTraffic :one
UPDATE traffic
SET traffic_date = $2,
avgdailydldatamb = $3,
avgdailyuldatamb = $4,
avgdailytotdatamb = $5,
avgdailytotvoicemin = $6,
avgdailytotvideomin = $7,
qci1_data = $8,
qci6_data = $9,
qci8_data = $10,
qci_other_data = $11,
avgdailytotvoicemin4g = $12,
avgdailytotvoicemintotal = $13,
userdlthroughput = $14,
dlpacketlossrate = $15,
overallpsdropcallrate = $16,
bhdldatamb = $17,
bhupdatamb = $18,
bhtotdatamb = $19,
bhtotvoicemin = $20,
bhtotvideomin = $21,
bhcsusers = $22,
bhhsupausers = $23,
bhhsdpausers = $24,
bhr99uldl = $25,
powercapacity = $26,
powerutilization = $27,
codecapacity = $28,
codeutilization = $29,
ceulcapacity = $30,
ceulutilization = $31,
cedlcapacity = $32,
cedlutilization = $33,
iubcapacity = $34,
iubutlization = $35,
bhrrcusers = $36,
cell_id = $37
WHERE id = $1
RETURNING *;

-- name: DeleteTraffic :exec
DELETE FROM traffic
WHERE id = $1;
