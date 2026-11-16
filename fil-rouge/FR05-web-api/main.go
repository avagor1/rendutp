package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Meteo struct {
	Current struct {
		Time        string  `json:"time"`
		Temperature float64 `json:"temperature_2m"`
		Vent        float64 `json:"wind_speed_10m"`
	} `json:"current"`
}

func recupererMeteo(baseURL string, lat, lon float64) (Meteo, error) {
	var m Meteo

	adresse := fmt.Sprintf(
		"%s/v1/forecast?latitude=%v&longitude=%v&current=temperature_2m,wind_speed_10m",
		baseURL, lat, lon)

	resp, err := http.Get(adresse)
	if err != nil {
		return m, fmt.Errorf("appel impossible : %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return m, fmt.Errorf("l'API a répondu %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		return m, fmt.Errorf("JSON illisible : %w", err)
	}

	return m, nil
}

func main() {
	m, err := recupererMeteo("http://localhost:8081", 48.89, 2.20)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Météo du campus à %s : %.1f °C, vent %.1f km/h\n",
		m.Current.Time, m.Current.Temperature, m.Current.Vent)
}
