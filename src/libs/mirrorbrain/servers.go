package mirrorbrain

import (
	"libs/database"
)

type MirrorbrainServer struct {
	database.Server
	Distance int
}
type MirrorbrainServers []MirrorbrainServer

func ChooseServer(requestFile RequestFile, requestIp string, servers database.Servers) database.Server {
	geoInfo := GeoLookup(requestIp)
	mirrorbrainServers := filterServers(requestFile, servers)
	prepareServerLists(requestFile, geoInfo, mirrorbrainServers)
	return servers[0]
}

func filterServers(requestFile RequestFile, servers database.Servers) (mirrorbrainServers MirrorbrainServers) {
	for _, server := range servers {
		//  skip mirror because of server.FileMaxsize
		if server.FileMaxsize >= requestFile.StatInfo.Size() {
			var mbServer = MirrorbrainServer{server, 0}
			mbServer.Distance = 123
			//  calculate distance: (int) ( sqrt( pow((lat - new->lat), 2) + pow((lng - new->lng), 2) ) * 1000 );
			mirrorbrainServers = append(mirrorbrainServers, mbServer)
		}
	}
	return mirrorbrainServers
}

// servers should not contain illegal entries (i.e. null values)
func prepareServerLists(requestFile RequestFile, geoInfo *GeoInfo, servers MirrorbrainServers) {
	// iterate all servers and build lists

	// TODO use search prio to exit early so redirect search doesn't build all lists
	//  list: all
	//  list: server with same prefix
	//	  samePrefix := make([]database.Servers, 0, len(servers))
	//  list: server in same AS
	//  list: server in same country (wildcard country codes)
	//  list: fallback country codes)
	//  list: same region (allows foreign country)
	//  list: server elsewhere (allows foreign regions)
	// fallback logic: same country empty? use fallback country
	// fallback: empty list of all servers use cfg->fallbacks
	// if geoip: best_method = closest
	// else best_method = rank
	// sort lists by method (same_prefix, same_as, same_country, same_region, elsewhere)

}

// func ServersInAsn(servers database.Servers, asn string) database.Servers {

// }

// func ServersInSameCountry(servers database.Servers, string country) database.Servers {
// }
// func ServersOnSameContinent(servers database.Servers) database.Servers {
// }
// func ServersOther(servers database.Servers) database.Servers {
// }

func maxScore(servers database.Servers) database.Server {
	max := servers[0]

	for _, server := range servers {
		if server.Score > max.Score {
			max = server
		}
	}
	return max
}
