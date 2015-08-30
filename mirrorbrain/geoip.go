package mirrorbrain

import (
	"github.com/manno/geoip"
	"github.com/manno/mirrorbrain-server/database"
	"log"
)

var file = "/usr/share/GeoIP/GeoIPCity.dat"
var geoIP *geoip.GeoIP

type GeoInfo struct {
	RequestIp     string
	AS            string
	Prefix        string
	CountryName   string
	CountryCode   string
	ContinentCode string
	Latitude      float64
	Longitude     float64
	RegionCode    string
	RegionName    string
}

func (geoInfo GeoInfo) HasCoords() bool {
	return geoInfo.Latitude != 0 && geoInfo.Longitude != 0
}

func GeoSetup() {
	gi, err := geoip.Open(file)
	if err != nil {
		log.Fatal("Could not open GeoIP database\n")
	}
	geoIP = gi
}

// TODO support ipv6 address lookups
func GeoLookup(ip string) *GeoInfo {
	geoInfo := new(GeoInfo)
	geoInfo.RequestIp = ip

	// lookup autonomous system number for request ip
	pfx, as, err := database.SelectAsn(ip)
	if err != nil {
		geoInfo.AS = as
		geoInfo.Prefix = pfx
	}

	// lookup geo information for ip
	geoIpRecord := geoIP.GetRecord(ip)
	if geoIpRecord == nil {
		log.Println("Failed to acquire GeoIP record for:", ip)
		return geoInfo
	}

	geoInfo.CountryName = fetch(geoIpRecord.CountryName)
	geoInfo.CountryCode = fetch(geoIpRecord.CountryCode)
	geoInfo.CountryName = fetch(geoIpRecord.CountryName)
	geoInfo.ContinentCode = fetch(geoIpRecord.ContinentCode)
	geoInfo.Latitude = float64(geoIpRecord.Latitude)
	geoInfo.Longitude = float64(geoIpRecord.Longitude)
	geoInfo.RegionCode = fetch(geoIpRecord.Region)
	//geoInfo.RegionName = geoIP.GetRegionName(geoIpRecord.CountryCode, geoIpRecord.Region)

	return geoInfo
}

// convert empty values to -- because mirrorbrain does this, too
func fetch(str string) string {
	if str == nil {
		return "--"
	}
	return str
}
