package database

import (
	"database/sql"
	"time"
)

type ResultRow struct {
	Id        int
	ApiKey    string
	Ip        sql.NullString
	OldIp     sql.NullString
	UpdatedAt time.Time
}

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
                           WHERE path = %s)::smallint[]) 
                      AND enabled AND status_baseurl AND score > 0`

const hashQuery = `SELECT file_id, md5hex, sha1hex, sha256hex, 
                          sha1piecesize, sha1pieceshex, btihhex, pgp, 
			  zblocksize, zhashlens, zsumshex 
			FROM hexhash 
			WHERE file_id = (SELECT id FROM filearr WHERE path = %s) 
			AND size = %lld
			AND mtime = %lld
			LIMIT 1`

/*

func FindUser(api_key string) (user models.User, err error) {
	rows, err := db.Query(findByApiKeyStmt, api_key)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if !rows.Next() {
		log.Printf("not found, params: (api_key: %s)", api_key)
		return user, errors.New("record not found")
	}
	user, err = scanUserRow(rows)
	if err != nil {
		log.Fatal(err)
	}

	return user, nil
}

func AllUsers() (users models.Users, err error) {
	rows, err := db.Query(allUsers)
	if err != nil {
		log.Printf("%s)", err)
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		user, err := scanUserRow(rows)
		if err != nil {
			log.Fatal("Failed to scan")
		}
		users = append(users, user)
	}

	return users, err
}

func ChangedUsers() (users models.Users, err error) {
	rows, err := db.Query(changedUsers)
	if err != nil {
		log.Printf("%s)", err)
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		user, err := scanUserRow(rows)
		if err != nil {
			log.Fatal("Failed to scan")
		}
		users = append(users, user)
	}

	return users, err
}

func UpdateUser(user models.User) {
	stmt, err := db.Prepare(updateUserStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Ip, user.OldIp, user.UpdatedAt, user.ApiKey); err != nil {
		log.Fatal(err)
	}
}

func UpdateUserLastChecked(user models.User) {
	stmt, err := db.Prepare(updateUserLastCheckedStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.ApiKey); err != nil {
		log.Fatal(err)
	}
}

func scanUserRow(rows *sql.Rows) (user models.User, err error) {
	var u UserRow
	if err := rows.Scan(&u.ApiKey, &u.Ip, &u.OldIp, &u.UpdatedAt); err != nil {
		return user, err
	}

	user.ApiKey = u.ApiKey
	user.Ip = convertSqlString(u.Ip)
	user.OldIp = convertSqlString(u.OldIp)
	user.UpdatedAt = u.UpdatedAt
	return user, err
}
*/
