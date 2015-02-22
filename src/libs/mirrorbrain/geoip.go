package mirrorbrain

import (
	"github.com/abh/geoip"
	"log"
)

var file = "/usr/share/GeoIP/GeoIP.dat"
var geoIP *geoip.GeoIP

func GeoSetup() {
	gi, err := geoip.Open(file)
	if err != nil {
		log.Fatal("Could not open GeoIP database\n")
	}
	geoIP = gi
}

// TODO country name
//  continent code
//  lat long
//  region region name
// TODO ipv6
func GeoLookup(ip string) (string, string) {
	country, prefix := geoIP.GetCountry(ip)
	return country, prefix
}
