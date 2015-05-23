package mirrorbrain

import "testing"

func TestLookup(t *testing.T) {
	GeoSetup()
	geoInfo := GeoLookup("178.63.212.179")
	if geoInfo.CountryCode != "DE" {
		t.Error("Expected to country to be DE, but was:", geoInfo.CountryCode)
	}
	geoInfo = GeoLookup("74.208.4.167")
	if geoInfo.CountryCode != "US" {
		t.Error("Expected to country to be US, but was:", geoInfo.CountryCode)
	}
}
