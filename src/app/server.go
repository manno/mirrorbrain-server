package main

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

	// TODO handle methods: head get propfind options?

	log.Println(r.RemoteAddr, requestPath, absPath)
	// TODO handle meta formats
	if isExtension("meta4", requestPath) {
		log.Println("can't handle meta4")
		http.Error(w, "not implemented yet", http.StatusInternalServerError)
	} else if isExtension("torrent", requestPath) {
		log.Println("can't handle torrent")
		http.Error(w, "not implemented yet", http.StatusInternalServerError)
	} else if isExtension("mirrorlist", requestPath) {
		log.Println("can't handle mirrorlist")
		http.Error(w, "not implemented yet", http.StatusInternalServerError)
	} else if isDir(absPath) {
		log.Println("Is dir!")
		// TODO dir listing
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
