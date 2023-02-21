package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/exec"
	"testing"
)

const dbDriver string = "postgres"
const dbSource string = "postgresql:///ratings?sslmode=disable"
const dbName string = "test_ratings"

var db *sql.DB

var TestQueries *Queries

func TestMain(m *testing.M) {
	CreateDB()
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("could not establish connection", err)
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
