package main

import (
	"flag"
	"fmt"
	"log"
	"weather/geo"
	"weather/weather"
)

func main() {
	city := flag.String("city", "", "Город")
	format := flag.Int("format", 1, "Формат вывода")
	flag.Parse()

	fmt.Println(*city)
	fmt.Println(*format)

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(geoData)
	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
