package main

import (
	"libs/database"
	"libs/mirrorbrain"
	"log"
	"net/http"
)

func sendRedirect(w http.ResponseWriter, r *http.Request, requestFile mirrorbrain.RequestFile) {
	// TODO decline redirect for:
	//   small files cfg->min_sie
	//   excluded files cfg->exclude_filemask
	//   excluded source ips cfg->exclude_ips
	//   excluded network cfg->exclude_networks
	//   excluded mime type cfg->exclude_mime
	//   excluded user agent cfg->exclude_agents
	// TODO since we don't send directly, what happens?

	path := requestFile.Path()
	servers, err := database.FindServers(path)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	dumpServers(servers)
	serverSelection := mirrorbrain.ChooseServer(requestFile, r.RemoteAddr, servers)
	http.Redirect(w, r, serverSelection.Chosen.RedirectUrl(path), http.StatusFound)
}

func dumpServers(servers database.Servers) {
	for _, server := range servers {
		log.Printf("%d %s (%s)\n", server.Score, server.Identifier, server.BaseUrl)
	}
}
