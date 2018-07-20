package main

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestGenerateGeohash(t *testing.T) {
  lat := 12.941084
  lon := 77.6099103
  expectedGeohash := "tdr1w5"
  actualGeohash := generateGeohash(lat, lon)
  assert.Equal(t, expectedGeohash, actualGeohash, "Generated geohash doesnt match with expected geohash.")
}

func TestParseLatLonFromMessage1(t *testing.T) {
  message := "12.941084,77.6099103"
  expectedLat, expectedLon := 12.941084, 77.6099103
  actualLat, actualLon := parseLatLonFromMessage(message)
  assert.Equal(t, expectedLat, actualLat, "Generated latitude doesnt match with expected latitude.")
  assert.Equal(t, expectedLon, actualLon, "Generated longitude doesnt match with expected longitude.")
}

func TestParseLatLonFromMessage2(t *testing.T) {
  message := "40.744630, -73.981481"
  expectedLat, expectedLon := 40.744630, -73.981481
  actualLat, actualLon := parseLatLonFromMessage(message)
  assert.Equal(t, expectedLat, actualLat, "Generated latitude doesnt match with expected latitude.")
  assert.Equal(t, expectedLon, actualLon, "Generated longitude doesnt match with expected longitude.")
}

func TestGetCountryFromGeohashToCountryMapping(t *testing.T) {
  geohash6FromCountryAUS, expectedCountryForAUSGeohash := "qsxtqw", "AUS"
  geohash6FromCountryJPN, expectedCountryForJPNGeohash := "xn7735", "JPN"
  geohash6FromOcean, expectedCountryForOceanGeohash := "mxty6f", ""

  geohashToCountryMapping := make(map[string]string)
  geohashToCountryMapping["qsxtqw"] = "AUS"
  geohashToCountryMapping["xn7"] = "JPN"

  actualCountryForAUSGeohash := getCountryFromGeohashToCountryMapping(geohash6FromCountryAUS, geohashToCountryMapping)
  actualCountryForJPNGeohash := getCountryFromGeohashToCountryMapping(geohash6FromCountryJPN, geohashToCountryMapping)
  actualCountryForOceanGeohash := getCountryFromGeohashToCountryMapping(geohash6FromOcean, geohashToCountryMapping)

  assert.Equal(t, expectedCountryForAUSGeohash, actualCountryForAUSGeohash, "Generated country doesnt match with expected country.")
  assert.Equal(t, expectedCountryForJPNGeohash, actualCountryForJPNGeohash, "Generated country doesnt match with expected country.")
  assert.Equal(t, expectedCountryForOceanGeohash, actualCountryForOceanGeohash, "Generated country doesnt match with expected country.")
}

func TestGetGeohashToCountryMapping(t *testing.T) {
  expectedGeohashToCountryMapping := make(map[string]string)
  expectedGeohashToCountryMapping["d6nq9t"] = "ABW"
  expectedGeohashToCountryMapping["wh3x0u"] = "IND"

  actualGeohashToCountryMapping := getGeohashToCountryMapping("dummy_test_data.csv.gz")

  assert.Equal(t, expectedGeohashToCountryMapping, actualGeohashToCountryMapping, "Generated geohashToCountryMapping doesnt match with expected geohashToCountryMapping.")
}
