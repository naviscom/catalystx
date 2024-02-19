package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/naviscom/catalystx/util"
	//_ "github.com/lib/pq"
)

//const (
//dbDriver = "postgres"
//dbSource = "postgresql://root:secret@localhost:5432/catalyst?sslmode=disable"
//)

// var testQueries *Queries
// var testDB *sql.DB
var testStore *Store

func TestMain(m *testing.M) {
	//var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
