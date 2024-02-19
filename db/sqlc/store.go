package db

import (
	//"database/sql"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store interface provides all functions to execute SQL queries and transactions
//type Store interface {
//Querier
//}

// SQLStore provides all functions to execute SQL queries and transactions
type Store struct {
	connpool *pgxpool.Pool
	*Queries
}

// NewStore creates a new store
func NewStore(connpool *pgxpool.Pool) *Store {
	return &Store{
		connpool: connpool,
		Queries:  New(connpool),
	}
}
