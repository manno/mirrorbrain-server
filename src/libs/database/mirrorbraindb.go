package database

import (
	"database/sql"
	"log"
)

type Server struct {
	Id             int
	Identifier     string
	Reqion         string
	Country        string
	Lat            float32
	Lng            float32
	Asn            int
	Prefix         string
	Score          int
	BaseUrl        string
	RegionOnly     bool
	CountryOnly    bool
	ASOnly         bool
	PrefixOnly     bool
	OtherCountries string
	FileMaxsize    int
}

type Servers []Server

const defaultQuery = `SELECT id, identifier, region, country,
                             lat, lng, 
                             asn, prefix, score, baseurl, 
                             region_only, country_only, 
                             as_only, prefix_only, 
                             other_countries, file_maxsize 
                      FROM server 
                      WHERE id::smallint = any(
                          (SELECT mirrors 
                           FROM filearr 
                           WHERE path = $1)::smallint[]) 
                      AND enabled AND status_baseurl AND score > 0`

const hashQuery = `SELECT file_id, md5hex, sha1hex, sha256hex, 
                          sha1piecesize, sha1pieceshex, btihhex, pgp, 
			  zblocksize, zhashlens, zsumshex 
			FROM hexhash 
			WHERE file_id = (SELECT id FROM filearr WHERE path = $1) 
			AND size = $2
			AND mtime = $3
			LIMIT 1`

func FindServers(path string) (servers Servers, err error) {
	rows, err := db.Query(defaultQuery, path)
	if err != nil {
		log.Printf("%s", err)
		return servers, err
	}
	defer rows.Close()

	for rows.Next() {
		server, err := scanServerRow(rows)
		if err != nil {
			log.Fatal("Failed to scan", err)
		}
		servers = append(servers, server)
	}

	return servers, err
}

func scanServerRow(rows *sql.Rows) (server Server, err error) {
	if err := rows.Scan(
		&server.Id,
		&server.Identifier,
		&server.Reqion,
		&server.Country,
		&server.Lat,
		&server.Lng,
		&server.Asn,
		&server.Prefix,
		&server.Score,
		&server.BaseUrl,
		&server.RegionOnly,
		&server.CountryOnly,
		&server.ASOnly,
		&server.PrefixOnly,
		&server.OtherCountries,
		&server.FileMaxsize); err != nil {
		return server, err
	}

	return server, err
}
