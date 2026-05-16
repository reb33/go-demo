package weather_test

import (
	"strings"
	"testing"
	"weather/geo"
	"weather/weather"
)

func TestGetWeather(t *testing.T) {
	city := "123123"
	geoData := geo.GeoData{
		City: city,
	}
	
	got, err := weather.GetWeather(geoData, 3)

	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	if got == "" {
		t.Fatal("Got empty string")
	}

	if !strings.Contains(got, strings.ToLower(city)){
		t.Fatalf("Ожидалось: %v, получено: %v", city, got)
	}
}

var testCases = []struct {
	name string
	city string
}{
	{name: "NotExistCity", city: "CityNotExist"},
	{name: "CityWithSpace", city: "Moscow "},
}

func TestGetWeatherCityNotExist(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			geoData := geo.GeoData{
				City: tc.city,
			}
			expected := "location not found: location not found\n"

			got, err := weather.GetWeather(geoData, 1)

			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if got != expected {
				t.Errorf("Ожидалось: %v, получено: %v", expected, got)
			}
		})
	}
}
