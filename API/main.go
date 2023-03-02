package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

// constants
const apiKey = "c500ece893c7c3ec9423ec5d8e5da39a"

// Structs
type WData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

type ABC struct {
	Val []string `json:"cities"`
}
type ResData struct {
	ActualTemperature float64  `json:"actual_temperature"`
	DiffTemperatures  []string `json:"diff_temperatures"`
}

// Methods
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {

	//Getting data from params....
	params := r.URL.Query()
	var cities []string

	for _, city := range params {
		cities = append(cities, city[0])
	}
	fmt.Println(cities)

	// body, _ := ioutil.ReadAll(r.Body)

	// var cities ABC
	// err := json.Unmarshal(body, &cities)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//	cities := []string{"amritsar", "delhi", "chennai", "london"}
	temperature := make(map[string]float64)

	FData := make(map[string]ResData)

	//Api calls based on cities....
	for _, city := range cities {
		fmt.Println("City name is: ", city)
		apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
		res, err := http.Get(apiUrl)
		if err != nil {
			http.Error(w, "error querrying the url ", http.StatusInternalServerError)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		var weatherdata WData
		err = json.Unmarshal(body, &weatherdata)
		temperature[city] = weatherdata.Main.Temp
	}

	fmt.Println("Temperature is: ", temperature)

	//Calculting diff of each pair of temperatures
	for i := range cities {
		var diffTemperatures []string

		for j := 0; j < len(cities); j++ {
			if i == j {
				continue
			}
			diff := temperature[cities[i]] - temperature[cities[j]]
			var key string
			if diff < 0 {
				key = fmt.Sprintf("%s is %v K colder than %s", cities[i], roundFloat(diff, 2), cities[j])
			} else {
				key = fmt.Sprintf("%s is  %v K hoter than %s", cities[i], roundFloat(diff, 2), cities[j])
			}
			diffTemperatures = append(diffTemperatures, key)
		}
		FData[cities[i]] = ResData{
			ActualTemperature: temperature[cities[i]],
			DiffTemperatures:  diffTemperatures,
		}
	}

	//converting to json Format
	resJSON, err := json.Marshal(FData)
	if err != nil {
		http.Error(w, "error encoding response as JSON", http.StatusInternalServerError)
		return
	}
	fmt.Println("Res Data is:", string(resJSON))
	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)

}

// Maon method
func main() {
	fmt.Println("Listening on port 8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/api", ApiHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))

}
