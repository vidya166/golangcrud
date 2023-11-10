package main

import (
	"encoding/json"
	_ "encoding/json"
	_ "fmt"
	"io/ioutil"
	"log"
	_ "math/rand"
	"net/http"
	_ "net/http"
	_ "strconv"
)

type Wthr struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Rain struct {
		OneH float64 `json:"1h"`
	} `json:"rain"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	req, err := http.NewRequest(http.MethodGet, "https://api.openweathermap.org/data/2.5/weather?lat=44.34&lon=10.99&appid=f2b60e478e7bb414ad2b1a28a4cb148d",
		nil)
	if err != nil {
		log.Fatalf("Failed to create request object for /GET endpoint: %v", err)
	}
	req.Header.Add("Content-type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send HTTP request: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	defer resp.Body.Close()
	var data Wthr
	json.Unmarshal(body, &data)

	json.NewEncoder(w).Encode(data)
}
