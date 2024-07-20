package test_dao

import (
	"database/sql"
	"log"
	"os"
	"simple-bank/src/dao"
	api "simple-bank/src/service"
	"testing"
	"time"

	"github.com/containerd/containerd/api/services/containers/v1"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	containers "github.com/testcontainers/testcontainers-go"
)

const (
	dbDriver = "postgres"
	dbName   = "simple_bank"
	dbUser   = "root"
	dbPwd    = "secret"
	// dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var (
	testQueries *dao.Queries
	testDB      *sql.DB
)

type TestDatabase struct {
	DbInstance *pgxpool.Pool
	DbAdress   string
	container  containers.container
}

func createContainer() (containers.Container, *pgxpool.Pool, string) {

}

func NewTestServer(store dao.Store) (*api.Server, error) {
	config := api.Config{
		TokenSymmetricKey:   "12345678912345678912345678912345",
		AccessTokenDuration: time.Minute,
	}

	return api.NewServer(config, store)
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the database")
	}
	testQueries = dao.New(testDB)

	os.Exit(m.Run())
}
