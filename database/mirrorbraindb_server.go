package database

import (
	"database/sql"
	"log"
)

type Server struct {
	Id             int
	Identifier     string
	Region         string
	Country        string
	Lat            float64
	Lng            float64
	Asn            string
	Prefix         string
	Score          int
	BaseUrl        string
	RegionOnly     bool
	CountryOnly    bool
	ASOnly         bool
	PrefixOnly     bool
	OtherCountries string
	FileMaxsize    int64
}

func (server Server) Valid() bool {
	if server.Score == nil {
		return false
	}
	if server.BaseUrl == nil || server.BaseUrl == "" {
		return false
	}
	if server.FileMaxsize == nil {
		return false
	}
	return true
}

type Servers []Server

const serversQuery = `SELECT id, identifier, region, country,
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

func FindServers(path string) (servers Servers, err error) {
	rows, err := db.Query(serversQuery, path)
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
		if server.Valid {
			servers = append(servers, server)
		}
	}

	return servers, err
}

func scanServerRow(rows *sql.Rows) (server Server, err error) {
	if err := rows.Scan(
		&server.Id,
		&server.Identifier,
		&server.Region,
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
