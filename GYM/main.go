package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Members struct {
	ID           int     `json:"id"`
	StartDate    string  `json:"start_date"`
	Name         string  `json:"name"`
	MemberShip   string  `json:"member_ship"`
	TotalPrice   float64 `json:"total_price"`
	EndDate      string  `json:"end_date"`
	MonthlyPrice float64 `json:"monthly_price"`
	Duration     int     `json:"duration"`
}

type Exit struct {
	Duration float64 `json:"duration"`
	Refund   float64 `json:"refund"`
}

var memType = map[string]float64{
	"Silver": 2000,
	"Gold":   1000,
}

var memberList = []Members{
	{
		StartDate:    "2023-02-02",
		ID:           0,
		Name:         "John",
		MemberShip:   "Silver",
		TotalPrice:   12000,
		EndDate:      "2023-03-08",
		MonthlyPrice: 2000,
		Duration:     6,
	},
	{
		StartDate:    "2023-03-02",
		ID:           1,
		Name:         "Johny",
		MemberShip:   "Gold",
		TotalPrice:   6000,
		EndDate:      "2023-03-08",
		MonthlyPrice: 1000,
		Duration:     6,
	},
}

func EnrollHandler(w http.ResponseWriter, r *http.Request) {
	var member Members
	err := json.NewDecoder(r.Body).Decode(&member)
	if err != nil {
		panic(err)
	}
	member.ID = len(memberList) + 1
	dateStr := time.Now().Truncate(time.Hour)
	member.StartDate = dateStr.Format("2006-01-02")

	if member.MemberShip == "Gold" {
		member.MonthlyPrice = memType[member.MemberShip]
		member.TotalPrice = float64(member.Duration) * member.MonthlyPrice
	} else if member.MemberShip == "Silver" {
		member.MonthlyPrice = memType[member.MemberShip]
		member.TotalPrice = float64(member.Duration) * member.MonthlyPrice
	} else {
		http.Error(w, "Please add valid membership", http.StatusBadRequest)
	}

	member.EndDate = dateStr.AddDate(0, 0, member.Duration*30).Format("2006-01-02")
	fmt.Println("member is", member)

	memberList = append(memberList, member)

	resJSON, err := json.Marshal(member)
	if err != nil {
		panic(err)
	}
	w.Write(resJSON)

}

func MemberHandler(w http.ResponseWriter, r *http.Request) {
	resJSON, err := json.Marshal(memberList)
	if err != nil {
		panic(err)
	}
	w.Write(resJSON)
}

func PriceGetHandler(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(memType)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(data))

}

func PriceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]float64
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	if goldPrice, ok := data["Gold"]; ok && goldPrice != 0 {
		memType["Gold"] = goldPrice
	}
	if silverPrice, ok := data["Silver"]; ok && silverPrice != 0 {
		memType["Silver"] = silverPrice
	}
	resJSON, err := json.Marshal(memType)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(resJSON))

}

func EndMemberShipHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query().Get("id")
	//refund
	now := time.Now().Truncate(24 * time.Hour)

	for _, member := range memberList {
		id, _ := strconv.Atoi(params)
		if id == member.ID {
			startDate, err := time.Parse("2006-01-02", member.StartDate)
			if err != nil {
				log.Fatal(err)
			}
			duration := float64(now.Sub(startDate).Hours() / 24)
			fmt.Println("Duration is", duration)
			// oneDayMoney := math.Round(member.TotalPrice / (float64(member.Duration) * 30))

			// MoneyRefund := (50 * (member.TotalPrice - (duration * oneDayMoney))) / 100

			// res := Exit{
			// 	Duration: duration,
			// 	Refund:   MoneyRefund,
			// }
			// json.NewEncoder(w).Encode(res)

			// member.Duration = int(duration)
			// member.EndDate = time.Now().AddDate(0, 0, int(duration)).Format("2006-01-02")
			// member.TotalPrice -= MoneyRefund
			return
		}
	}
}
func main() {
	fmt.Println("Listening on posrt:8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/enroll", EnrollHandler)
	mux.HandleFunc("/member", MemberHandler)
	mux.HandleFunc("/PriceGet", PriceGetHandler)
	mux.HandleFunc("/PriceUpdate", PriceUpdateHandler)
	mux.HandleFunc("/end-memberShip", EndMemberShipHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))

}
