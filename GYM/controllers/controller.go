package cont

import (
	"encoding/json"
	"fmt"
	mod "gym-api/models"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"time"
)

var jwtKey = []byte("secret_key")
var memType = map[string]float64{
	"Silver": 2000,
	"Gold":   1000,
}

// declaring  memberList to store the members of gym
var memberList = []mod.Members{
	{
		StartDate:    "01 Feb 2023",
		ID:           1,
		Name:         "John",
		MemberShip:   "Silver",
		TotalPrice:   12000,
		EndDate:      "01 Aug 2023",
		MonthlyPrice: 2000,
		Duration:     6,
	},
	{
		StartDate:    "01 Feb 2023",
		ID:           2,
		Name:         "Johny",
		MemberShip:   "Gold",
		TotalPrice:   5500,
		EndDate:      "27 Aug 2023",
		MonthlyPrice: 1000,
		Duration:     5.5,
	},
}

// EnrollHandler handles the request for enrolling new members.
func EnrollHandler(w http.ResponseWriter, r *http.Request) {

	// Create an empty member struct to store data from request body.
	var member mod.Members

	//getting data from response body and storing it in memmbers struct
	err := json.NewDecoder(r.Body).Decode(&member)
	if err != nil {
		panic(err)
	}
	//genrating a unique id for the new member
	member.ID = len(memberList) + 1

	// Set the start date of the membership as the current date and time.

	dateStr := time.Now().Truncate(time.Hour)
	member.StartDate = dateStr.Format("02 Jan 2006")

	// Calculate the total price of the membership based on the duration and membership type.
	if member.MemberShip == "Gold" {
		member.MonthlyPrice = memType[member.MemberShip]
		member.TotalPrice = float64(member.Duration) * member.MonthlyPrice
	} else if member.MemberShip == "Silver" {
		member.MonthlyPrice = memType[member.MemberShip]
		member.TotalPrice = float64(member.Duration) * member.MonthlyPrice
	} else {
		http.Error(w, "Please add valid membership", http.StatusBadRequest)
	}

	//checking if the membership duratiion is less than one month or not
	if member.Duration <= 1 {
		w.Write([]byte("Duration must be greater than 1"))
		return
	}

	//member.Balance = member.MoneySubmitted - member.TotalPrice

	// Set the end date of the membership based on the duration and current date.
	member.EndDate = dateStr.AddDate(0, 0, int(member.Duration*30)).Format("02 Jan 2006")
	fmt.Println("member is", member)

	// Add the new member to the member list.
	memberList = append(memberList, member)

	// Convert the new member data to JSON and write it to the response.
	resJSON, err := json.Marshal(member)
	if err != nil {
		panic(err)
	}
	w.Write(resJSON)

}

// MemberHandler gives the list of members present in a gym
func MemberHandler(w http.ResponseWriter, r *http.Request) {

	//sorting memberList based on their ID's
	sort.Slice(memberList, func(i, j int) bool {
		return memberList[i].ID < memberList[j].ID
	})
	//coverting the memberList to json and printing it out
	resJSON, err := json.Marshal(memberList)
	if err != nil {
		panic(err)
	}
	w.Write(resJSON)

}

// PriceGetHandler prints out the currents price of type of memberhip present in a gym
func PriceGetHandler(w http.ResponseWriter, r *http.Request) {

	data, err := json.Marshal(memType)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(data))

}

// PriceUpdateHandler handle the request to update the price of type of membership of a gym
func PriceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]float64

	//Taking input from body
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	//checking if the value is not equal to  or empty if not then changing va;ues of gold and silver respectively
	if goldPrice, ok := data["Gold"]; ok && goldPrice != 0 {
		memType["Gold"] = goldPrice
	}
	if silverPrice, ok := data["Silver"]; ok && silverPrice != 0 {
		memType["Silver"] = silverPrice
	}

	//converting the memType to json and printing it out
	resJSON, err := json.Marshal(memType)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(resJSON))

}

// EndMemberHandler hadles the request to end membership for a given member
func EndMemberShipHandler(w http.ResponseWriter, r *http.Request) {

	//get the member id from the query parameters
	params := r.URL.Query().Get("id")

	//get the current time trunctaed to the nearest day
	now := time.Now().Truncate(24 * time.Hour)

	//Loop through the member list to find the given member id
	for i, member := range memberList {
		id, err := strconv.Atoi(params)
		if err != nil {
			log.Fatal(err)
		}
		//check if the id matches with the current member id
		if id == member.ID {

			//calculate the dduration in days
			startDate, err := time.Parse("02 Jan 2006", member.StartDate)
			if err != nil {
				log.Fatal(err)
			}
			temp := now.Sub(startDate).Hours() / 24
			duration := float64(temp)
			fmt.Println("Duration is", duration)
			if duration < 30 {
				http.Error(w, "Cannot end membership before one month", http.StatusBadRequest)
				return

			}
			//calculate the monbey to refund to the member
			oneDayMoney := (member.TotalPrice / (float64(member.Duration) * 30))
			MoneyRefund := math.Round((member.TotalPrice - (duration * oneDayMoney)) / 2)

			//Update members information
			member.Duration = temp / 30
			member.EndDate = time.Now().AddDate(0, 0, int(duration)).Format("02 Jan 2006")
			member.TotalPrice -= MoneyRefund
			member.IsDelete = true

			// Create the exit information to be returned in the response
			res := mod.Exit{
				Duration: temp / 30,
				Refund:   MoneyRefund,
			}

			// Remove the old member information from the list and append the updated member
			memberList = append(memberList[:i], memberList[i+1:]...)
			memberList = append(memberList, member)

			// Write the member information and exit information to the response as JSON
			//json.NewEncoder(w).Encode(member)
			json.NewEncoder(w).Encode(res)

			return
		}
	}
	http.Error(w, "No member with the given ID was found", http.StatusNotFound)

}
