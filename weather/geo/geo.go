package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulationResponse struct {
	Error bool `json:"error"`
}

var ErrNoCity = errors.New("Такого города нет")
var ErrNot200 = errors.New("bad status code")

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			return nil, ErrNoCity
		}
		return &GeoData{City: city}, nil
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://ipapi.co/json/", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/148.0.0.0 Safari/537.36")
	// req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	// req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	// req.Header.Set("accept-encoding", "gzip, deflate, br, zstd")
	req.Header.Set("cookie", "csrftoken=mTD1JryK1vykDDCCEBa7rtz1MD7yyVwC; cf_clearance=KMGwLPDMsiVmcUoJX9Fg0LtOGCe8irEqySXshf8k8.E-1778877207-1.2.1.1-FGD3O2WtI02o0XHDpdSTqcE0D9t6DvJbpm8bZs9DS4gMUhrdtSJDaCezA4dN8E.QIqAj9AoRiEzWJXnRyHBotTSPMTdqebFYGSep0yUjjruPx7rlhWoAXikGP41A2VfioqcpiIzgWPJo1kIOvRwvjLSj4PoT5_cVzSfrtz.4rmu6R6tijR1m6EfGAXf5pDwEiUJQ2odCopq1GcxaV0XbuEHDEFGlFSCNAndd0YfvCDCqeW7lxbEf.ZUOP57MoMOQZZ_aATmO036VbeWJsdWq1ST.SIf4SXqRkmyfxCJrRVj82044O.OHVpkJRnZsRSB9IXNFfm1SLBpdEWWCLAx7aY9_EXZHlXvwER_9UI7d0FmS7DqCHq9NugrKMcT1YLO8INF4Za3g0quHeDwB8qI2bM4jxj02WEHTXuikdFeTi2Q")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, ErrNot200
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}


func checkCity(city string) bool {
	reqBody, _ := json.Marshal(map[string]string{
		"city": city,
	})
	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var cityPopulation CityPopulationResponse
	json.Unmarshal(body, &cityPopulation)
	return !cityPopulation.Error
}