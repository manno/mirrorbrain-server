package main

import (
	"libs/database"
	"libs/mirrorbrain"
	"log"
	"net"
	"net/http"
)

func sendRedirect(w http.ResponseWriter, r *http.Request, requestFile mirrorbrain.RequestFile) {

	path := requestFile.Path()
	servers, err := database.FindServers(path)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	dumpServers(servers)
	serverSelection := mirrorbrain.ChooseServer(requestFile, ExtractIP(r.RemoteAddr), servers)
	if serverSelection.FoundIn != "" {
		log.Println("Using server:", serverSelection)
		http.Redirect(w, r, serverSelection.Chosen.RedirectUrl(path), http.StatusFound)
	} else {
		log.Println("No usable mirrors after classification for:", requestFile.Path())
		http.NotFound(w, r)
	}
}

func dumpServers(servers database.Servers) {
	for _, server := range servers {
		log.Printf("%d %s (%s)\n", server.Score, server.Identifier, server.BaseUrl)
	}
}
