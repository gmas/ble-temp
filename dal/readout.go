package dal

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Readout struct {
	Uuid  string
	Date  string
	Temp  float64
	Humid float64
}

func getDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./foo.sql3")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

//func NewReadout() *Readout {
//}

func NewReadoutFromJson(readoutJson []byte) (*Readout, error) {
	var readout Readout
	if err := json.Unmarshal(readoutJson, &readout); err != nil {
		return nil, err
	}
	return &readout, nil
}

func (sensorVal *Readout) Insert() error {
	db := getDb()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Prepare INSERT statement
	stmt, err := tx.Prepare(
		"INSERT INTO sensor_data(time, temperature, humidity) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Execute statement
	// TODO add uuid to table and insert
	stmt.Exec(sensorVal.Date, sensorVal.Temp, 0)
	if err != nil {
		log.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func createDb() {
	//sqlStmt := `
	//CREATE TABLE sensor_data
	//	(id INTEGER NOT NULL PRIMARY KEY,
	//	 time datetime,
	//	 temperature FLOAT,
	//	 humidity FLOAT);
	//	 `
	//_, err = db.Exec(sqlStmt)
	//if err != nil {
	//	log.Printf("%q: %s\n", err, sqlStmt)
	//	return
	//}
}
