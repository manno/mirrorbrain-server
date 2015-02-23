package mirrorbrain

import (
	"libs/database"
)

func ChooseServer(requestIp string, servers database.Servers) database.Server {
	geoInfo := GeoLookup(requestIp)

	prepareServerLists(geoInfo, servers)
	return servers[0]
}

// servers should not contain illegal entries (i.e. null values)
func prepareServerLists(geoInfo *GeoInfo, servers database.Servers) {
	//samePrefix := make([]database.Servers, 0, len(servers))
	// iterate all servers and build lists
	//  calculate distance: (int) ( sqrt( pow((lat - new->lat), 2) + pow((lng - new->lng), 2) ) * 1000 );
	//  skip mirror because of file size
	// TODO use search prio so redirect search doesn't build all lists
	//  list: all
	//  list: server with same prefix
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
