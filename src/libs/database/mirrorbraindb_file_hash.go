package database

import (
	"database/sql"
	"log"
)

type FileInfo struct {
	Id            int
	Md5Hex        string
	Sha1Hex       string
	Sha256Hex     string
	Sha1PieceSize float32
	Sha1PiecesHex string
	BtihHex       int
	Pgp           string
	ZblockSize    int
	ZhashLens     string
	ZsumsHex      bool
}

const fileHashQuery = `SELECT file_id, md5hex, sha1hex, sha256hex, 
                          sha1piecesize, sha1pieceshex, btihhex, pgp, 
			  zblocksize, zhashlens, zsumshex 
			FROM hexhash 
			WHERE file_id = (SELECT id FROM filearr WHERE path = $1) 
			AND size = $2
			AND mtime = $3
			LIMIT 1`

func SelectFileInfo(path string) (fileInfo *FileInfo, err error) {
	rows, err := db.Query(fileHashQuery, path)
	if err != nil {
		log.Printf("%s", err)
		return fileInfo, err
	}
	defer rows.Close()

	rows.Next()
	fileInfo, err = scanFileInfoRow(rows)
	if err != nil {
		log.Fatal("Failed to scan", err)
	}

	return fileInfo, err
}

func scanFileInfoRow(rows *sql.Rows) (fileInfo *FileInfo, err error) {
	if err = rows.Scan(
		&fileInfo.Id,
		&fileInfo.Md5Hex,
		&fileInfo.Sha1Hex,
		&fileInfo.Sha256Hex,
		&fileInfo.Sha1PieceSize,
		&fileInfo.Sha1PiecesHex,
		&fileInfo.BtihHex,
		&fileInfo.Pgp,
		&fileInfo.ZblockSize,
		&fileInfo.ZhashLens,
		&fileInfo.ZsumsHex); err != nil {
		return fileInfo, err
	}

	return fileInfo, err
}
