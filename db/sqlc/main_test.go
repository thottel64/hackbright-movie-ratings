package db

import (
	"database/sql"
	"log"
	"os"
	"os/exec"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const dbDriver string = "postgres"
const dbSource string = "postgresql:///test_ratings?sslmode=disable"
const dbName string = "test_ratings"

var db *sql.DB

var TestQueries *Queries

func TestMain(m *testing.M) {
	CreateDB()
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("could not establish connection", err)
	}
	mig, err := migrate.New("file://../migrations", dbSource)
	if err != nil {
		log.Fatalln("error:", err)
	}
	err = mig.Up()
	if err != nil {
		log.Fatalln("error:", err)
	}
	TestQueries = New(conn)
	os.Exit(m.Run())
}
func CreateDB() {
	dropDb := exec.Command("dropdb", "--if-exists", dbName)
	createDb := exec.Command("createDb", dbName)

	err := dropDb.Run()
	if err != nil {
		log.Fatal("unable to drop test db:", err)
	}

	err = createDb.Run()
	if err != nil {
		log.Fatal("Unable to create new test db:", err)
	}
}
