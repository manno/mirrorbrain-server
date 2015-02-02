package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"libs/settings"
	"log"
)

var db *sql.DB

func Connect(config *settings.Config) {
	db = connectDatabase(config)
}

func Close() {
	db.Close()
}

func Exists(tableName string) bool {
	existsStmt := ` 
	SELECT EXISTS (
		SELECT 1 
		FROM   pg_catalog.pg_class c
		JOIN   pg_catalog.pg_namespace n ON n.oid = c.relnamespace
		WHERE  n.nspname = 'public'
		AND    c.relname = '%s'
		AND    c.relkind = 'r'
	);
	`

	var found bool
	err := db.QueryRow(existsStmt, tableName).Scan(&found)
	if err != nil {
		return false
	}
	return found
}

func convertSqlString(nullStr sql.NullString) string {
	if nullStr.Valid {
		return nullStr.String
	}
	return ""
}

func connectDatabase(config *settings.Config) *sql.DB {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", config.DbUser,
		config.DbPassword, config.DbName, config.DbHost)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
