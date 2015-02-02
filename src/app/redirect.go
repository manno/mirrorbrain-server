package main

import (
	"libs/config"
	"libs/database"
	"log"
	"net/http"
	"os"
	"path"
)

const WebRoot = "/usr/share/pixmaps"

func viewHandler(w http.ResponseWriter, r *http.Request) {
	var requestPath = path.Join(WebRoot, r.URL.Path)
	log.Println(r.RemoteAddr, r.URL.Path, requestPath)
	if _, err := os.Stat(requestPath); err == nil {
		// TODO handle meta formats
		// pos = strings.LastIndex(path, "meta4")
		// if pos > -1 && pos == len(path) - 5
		sendRedirect(w, r, path.Clean(r.URL.Path))
	} else {
		http.NotFound(w, r)
	}
}

func sendRedirect(w http.ResponseWriter, r *http.Request, path string) {
	// TODO find server in database
	//fmt.Fprintf(w, "%s", path)
	http.Redirect(w, r, "http://otherhost"+path, http.StatusFound)
}

func main() {
	config := new(config.Config)
	config.Setup()

	database.Connect(config)
	defer database.Close()

	var port = ":" + config.BindPort
	log.Printf("Listening on %s", config.BindPort)
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
