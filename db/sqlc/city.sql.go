// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: city.sql

package db

import (
	"context"
)

const createCity = `-- name: CreateCity :one
INSERT INTO cities (
    city_name,
    city_desc,
    state_id
) VALUES (
 $1, $2, $3
)
RETURNING id, city_name, city_desc, state_id
`

type CreateCityParams struct {
	CityName string `json:"city_name"`
	CityDesc string `json:"city_desc"`
	StateID  int64  `json:"state_id"`
}

func (q *Queries) CreateCity(ctx context.Context, arg CreateCityParams) (City, error) {
	row := q.db.QueryRow(ctx, createCity, arg.CityName, arg.CityDesc, arg.StateID)
	var i City
	err := row.Scan(
		&i.ID,
		&i.CityName,
		&i.CityDesc,
		&i.StateID,
	)
	return i, err
}

const deleteCity = `-- name: DeleteCity :exec
DELETE FROM cities
WHERE id = $1
`

func (q *Queries) DeleteCity(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteCity, id)
	return err
}

const getCity0 = `-- name: GetCity0 :one
SELECT id, city_name, city_desc, state_id FROM cities
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCity0(ctx context.Context, id int64) (City, error) {
	row := q.db.QueryRow(ctx, getCity0, id)
	var i City
	err := row.Scan(
		&i.ID,
		&i.CityName,
		&i.CityDesc,
		&i.StateID,
	)
	return i, err
}

const getCity1 = `-- name: GetCity1 :one
SELECT id, city_name, city_desc, state_id FROM cities
WHERE city_name = $1 LIMIT 1
`

func (q *Queries) GetCity1(ctx context.Context, cityName string) (City, error) {
	row := q.db.QueryRow(ctx, getCity1, cityName)
	var i City
	err := row.Scan(
		&i.ID,
		&i.CityName,
		&i.CityDesc,
		&i.StateID,
	)
	return i, err
}

const listCities = `-- name: ListCities :many
SELECT id, city_name, city_desc, state_id FROM cities
WHERE state_id = $3
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListCitiesParams struct {
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
	StateID int64 `json:"state_id"`
}

func (q *Queries) ListCities(ctx context.Context, arg ListCitiesParams) ([]City, error) {
	rows, err := q.db.Query(ctx, listCities, arg.Limit, arg.Offset, arg.StateID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []City{}
	for rows.Next() {
		var i City
		if err := rows.Scan(
			&i.ID,
			&i.CityName,
			&i.CityDesc,
			&i.StateID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCity = `-- name: UpdateCity :one
UPDATE cities
SET city_name = $2,
city_desc = $3,
state_id = $4
WHERE id = $1
RETURNING id, city_name, city_desc, state_id
`

type UpdateCityParams struct {
	ID       int64  `json:"id"`
	CityName string `json:"city_name"`
	CityDesc string `json:"city_desc"`
	StateID  int64  `json:"state_id"`
}

func (q *Queries) UpdateCity(ctx context.Context, arg UpdateCityParams) (City, error) {
	row := q.db.QueryRow(ctx, updateCity,
		arg.ID,
		arg.CityName,
		arg.CityDesc,
		arg.StateID,
	)
	var i City
	err := row.Scan(
		&i.ID,
		&i.CityName,
		&i.CityDesc,
		&i.StateID,
	)
	return i, err
}