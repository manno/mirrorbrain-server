package main

// TODO handle methods: head get propfind options?
// TODO handle meta formats
// TODO better handling of scoring (shuffle?)
// TODO use https://github.com/abh/geoip

import (
	"libs/database"
	"libs/settings"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var config = new(settings.Config)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	absPath := webrootPath(config.WebRoot, r.URL.Path)
	requestPath := path.Clean(r.URL.Path)

	log.Println(r.RemoteAddr, requestPath, absPath)
	if isExtension("meta4", requestPath) {
		// TODO set content type header: "application/metalink4+xml; charset=UTF-8"
		printMeta4(w, requestPath)
	} else if isExtension("torrent", requestPath) {
		log.Println("can't handle torrent")
		http.Error(w, "not implemented yet", http.StatusInternalServerError)
	} else if isExtension("mirrorlist", requestPath) {
		// TODO sort mirrorlist?
		printMirrorList(w, requestPath)
	} else if isDir(absPath) {
		printDirectoryList(w, absPath, requestPath)
	} else if _, err := os.Stat(absPath); err == nil {
		sendRedirect(w, r, requestPath)
	} else {
		http.NotFound(w, r)
	}
}

// format request, i.e: ?meta4 or .torrent
func isExtension(extension string, requestPath string) bool {
	pos := strings.LastIndex(requestPath, extension)
	if pos > -1 && pos == len(requestPath)-len(extension) {
		if requestPath[pos-1] == '.' || requestPath[pos-1] == '?' {
			return true
		}
	}
	return false
}

func main() {
	config.Setup()

	database.Connect(config)
	defer database.Close()

	var port = ":" + config.BindPort
	log.Printf("Listening on %s", config.BindPort)
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
