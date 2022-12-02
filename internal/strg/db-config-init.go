package strg

import (
	"database/sql"
	"example.com/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

func InitDB(dbConfig *config.DBConfig) *sql.DB {
	driverName, dataSource := parse(dbConfig)
	db, err := sql.Open(driverName, dataSource)

	migrationUp(db, dbConfig)

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

func migrationUp(db *sql.DB, dbConfig *config.DBConfig) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("Database migration error. " + err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance(
		dbConfig.MigrationFilesURL,
		dbConfig.DBName,
		driver)
	if err != nil {
		log.Fatal("Database migration error. " + err.Error())
	}
	m.Up()

	_, dirty, _ := m.Version()
	if dirty {
		log.Fatal("Dirty migration in database")
	}
}
