package geo_test

import (
	"testing"
	"weather/geo"
)

func TestGetMyLocation(t *testing.T) {
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}

	got, err := geo.GetMyLocation(city)

	if err != nil {
		t.Fatal("Ошибка получения города", err)
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %s, получено %s", expected.City, got.City)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "NotExistCity"
	if _, err := geo.GetMyLocation(city); err != geo.ErrNoCity {
		t.Errorf("Ожидалась ошибка %v, получено %v", geo.ErrNoCity, err)
	}

}
