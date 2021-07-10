package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type ConfigDB struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func (conf *ConfigDB) NewDB() (*sql.DB, error) {
	postgresInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s",
		conf.host, conf.port, conf.user, conf.password, conf.dbname)
	db, err := sql.Open("postgres", postgresInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
