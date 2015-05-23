package main

import (
	"github.com/manno/mirrorbrain-server/database"
	"github.com/manno/mirrorbrain-server/mirrorbrain"
	"github.com/manno/mirrorbrain-server/settings"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var config = new(settings.Config)

func ExtractIP(remoteAddr string) string {
	ip, _, err := net.SplitHostPort(remoteAddr)
	if err == nil {
		return ip
	} else {
		return remoteAddr
	}
}

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

	if isDir(absPath) {
		printDirectoryList(w, absPath, requestPath)

	} else if stat, err := os.Stat(absPath); err == nil {
		requestFile := mirrorbrain.RequestFile{requestPath, stat}

		if requestFile.HasExtension("meta4") {
			printMeta4(w, r, requestFile)

		} else if requestFile.HasExtension("mirrorlist") {
			printMirrorList(w, r, requestFile)

		} else if requestFile.HasExtension("sha256") {
			http.Error(w, "not implemented", http.StatusInternalServerError)
		} else if requestFile.HasExtension("sha1") {
			http.Error(w, "not implemented", http.StatusInternalServerError)
		} else if requestFile.HasExtension("md5") {
			http.Error(w, "not implemented", http.StatusInternalServerError)
		} else if requestFile.HasExtension("torrent") {
			log.Println("can't handle torrent")
			http.Error(w, "not implemented", http.StatusInternalServerError)

		} else {
			sendRedirect(w, r, requestFile)
		}

	} else {
		http.NotFound(w, r)
	}
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
