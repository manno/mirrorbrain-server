package main

import (
	"libs/database"
	"libs/mirrorbrain"
	"libs/settings"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var config = new(settings.Config)

// We do not support all mirrorbrain output formats. Specifically .torrent is missing
//
// TODO The redirector does not support declining redirects like the original does and isn't able to send files directly.
//   * small files cfg->min_size and send file directly instead
//   * excluded files cfg->exclude_filemask
//   * excluded source ips cfg->exclude_ips
//   * excluded network cfg->exclude_networks
//   * excluded mime type cfg->exclude_mime
//   * excluded user agent cfg->exclude_agents
//
// TODO handle all http methods: head get propfind options, how are they used in mirrorbrain?
func viewHandler(w http.ResponseWriter, r *http.Request) {
	absPath := webrootPath(config.WebRoot, r.URL.Path)
	requestPath := path.Clean(r.URL.Path)

	log.Println(r.RemoteAddr, requestPath, absPath)
	if isExtension("meta4", requestPath) {
		printMeta4(w, requestPath)
	} else if isExtension("torrent", requestPath) {
		log.Println("can't handle torrent")
		http.Error(w, "not implemented yet", http.StatusInternalServerError)
	} else if isExtension("mirrorlist", requestPath) {
		printMirrorList(w, requestPath)
	} else if isDir(absPath) {
		printDirectoryList(w, absPath, requestPath)
	} else if stat, err := os.Stat(absPath); err == nil {
		requestFile := mirrorbrain.RequestFile{requestPath, stat}
		sendRedirect(w, r, requestFile)
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

	mirrorbrain.GeoSetup()
	database.Connect(config)
	defer database.Close()

	var port = ":" + config.BindPort
	log.Printf("Listening on %s", config.BindPort)
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
