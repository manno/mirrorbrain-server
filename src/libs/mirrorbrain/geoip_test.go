package mirrorbrain

import "testing"

func TestLookup(t *testing.T) {
	GeoSetup()
	country, _ := GeoLookup("178.63.212.179")
	if country != "DE" {
		t.Error("Expected to country to be DE, but was:", country)
	}
	country, _ = GeoLookup("74.208.4.167")
	if country != "US" {
		t.Error("Expected to country to be US, but was:", country)
	}

}
