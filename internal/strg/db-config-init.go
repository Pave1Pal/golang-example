package strg

import (
	"database/sql"
	"example.com/internal/config"
	_ "github.com/lib/pq"
	"log"
)

func InitDB(dbConfig *config.DBConfig) *sql.DB {
	driverName, dataSource := parse(dbConfig)
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		log.Fatal(err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func parse(dbConfig *config.DBConfig) (string, string) {
	driverName := dbConfig.DriverName
	dataSource :=
		"user=" + dbConfig.UserName +
			" password=" + dbConfig.Password +
			" host=" + dbConfig.Host +
			" dbname=" + dbConfig.DBName +
			" sslmode=" + dbConfig.SslMode
	return driverName, dataSource
}
