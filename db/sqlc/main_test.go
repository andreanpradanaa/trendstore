package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/andreanpradanaa/trendstore/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conPool, err := pgxpool.New(context.Background(), config.DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect db:", err)
	}

	testStore = NewStore(conPool)
	os.Exit(m.Run())
}
