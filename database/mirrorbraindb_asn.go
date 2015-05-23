package database

import "log"

// prefix 2 as number
const asnQuery = `SELECT pfx, asn FROM pfx2asn WHERE pfx >>= ip4r($1) ORDER BY ip4r_size(pfx) LIMIT 1`

func SelectAsn(prefix string) (pfx string, asn string, err error) {
	rows, err := db.Query(asnQuery, prefix)
	if err != nil {
		log.Printf("%s", err)
		return pfx, asn, err
	}
	defer rows.Close()

	if !rows.Next() {
		return "", "", rows.Err()
	}

	if err := rows.Scan(&pfx, &asn); err != nil {
		log.Fatal("Failed to scan", err)
	}

	return pfx, asn, err
}
