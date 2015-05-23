package main

import (
	"github.com/manno/mirrorbrain-server/database"
	"github.com/manno/mirrorbrain-server/mirrorbrain"
	"net/http"
)

type mirrorlist struct {
	Url                string //http://cdn.media.ccc.de/congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4
	FileName           string //31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4
	Path               string // /congress/2014/h264-hd/31c3-6264-de-en-Wir_beteiligen_uns_aktiv_an_den_Diskussionen_hd.mp4
	SizeMB             string //523M
	SizeBytes          string //234 bytes
	LastModified       string // Mon, 29 Dec 2014 14:13:19 GMT (Unix time: 1419862399)
	Sha256             string
	Sha1               string
	Md5                string
	RequestIp          string
	RequestCountry     string
	RequestCountryCode string
	RequestNetwork     string // 78.34.0.0/15
	RequestAS          string
}

// TODO sort mirrorlist?
func printMirrorList(w http.ResponseWriter, r *http.Request, requestFile mirrorbrain.RequestFile) {
	path := requestFile.Path()
	servers, err := database.FindServers(path)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	mirrorbrain.CreateServerLists(requestFile, ExtractIP(r.RemoteAddr), servers)
	fileInfo, err = database.SelectFileInfo(path)
}
