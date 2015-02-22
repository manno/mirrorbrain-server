package database

import (
	"database/sql"
	"log"
)

// prefix 2 as number
const defaultQuery = `SELECT asn FROM pfx2asn
                      WHERE pfx = $1
		      `

func SelectAsn(pfx string) (asn string, err error) {
	rows, err := db.Query(hashQuery, path)
	if err != nil {
		log.Printf("%s", err)
		return asn, err
	}
	defer rows.Close()

	rows.Next()
	if err := rows.Scan(&asn); err != nil {
		log.Fatal("Failed to scan", err)
	}

	return asn, err
}
