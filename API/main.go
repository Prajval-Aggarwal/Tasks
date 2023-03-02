package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"sort"
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
	City              string             `json:"city"`
	ActualTemperature float64            `json:"actual_temperature"`
	DiffTemperatures  map[string]float64 `json:"diff_temperatures"`
}

// Methods
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	// fmt.Println("body is:", string(body))
	var cities ABC
	err := json.Unmarshal(body, &cities)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("cities is:", cities.Val)

	//	cities := []string{"amritsar", "delhi", "chennai", "london"}
	temperature := make(map[string]float64)

	var resData []ResData

	for _, city := range cities.Val {
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

	for i := range cities.Val {
		diffTemperatures := make(map[string]float64)

		for j := 0; j < len(cities.Val); j++ {
			if i == j {
				continue
			}
			diff := temperature[cities.Val[i]] - temperature[cities.Val[j]]
			//fmt.Printf("Temperature difference between %s and %s is: %.2f\n", cities.Val[i], cities.Val[j], diff)
			key := fmt.Sprintf("%s-%s", cities.Val[i], cities.Val[j])
			//	fmt.Println("Difference is: ", diff)
			diffTemperatures[key] = roundFloat(diff, 2)
		}
		resData = append(resData, ResData{
			City:              cities.Val[i],
			ActualTemperature: temperature[cities.Val[i]],
			DiffTemperatures:  diffTemperatures,
		})
	}

	//sorting done here
	sort.Slice(resData, func(i, j int) bool {
		return resData[i].ActualTemperature > resData[j].ActualTemperature
	})

	//converting to json Format
	resJSON, err := json.Marshal(resData)
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
