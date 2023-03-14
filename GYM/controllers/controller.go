package cont

import (
	"encoding/json"
	"fmt"
	mod "gym-api/models"
	"log"
	"math"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	fmt.Println("Connecting to database...")

	dsn := "host=localhost port=5432 user=postgres password=Test@123 dbname=gym sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in conncting to database", err)
	}
	db.AutoMigrate(&mod.Member{}, &mod.MemberShip{})
	DB = db

	fmt.Println("Connected to database....")
}

// EnrollHandler handles the request for enrolling new members.
func EnrollHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var member mod.Member

	err := json.NewDecoder(r.Body).Decode(&member)
	if err != nil {
		panic(err)
	}

	dateStr := time.Now().Truncate(time.Hour)
	member.StartDate = dateStr.Format("02 Jan 2006")

	var memShip mod.MemberShip
	DB.Where("name= ?", member.MemberShip_Name).First(&memShip)

	member.MonthlyPrice = memShip.Price
	member.TotalPrice = float64(member.Duration) * member.MonthlyPrice

	if member.Duration <= 1 {
		w.Write([]byte("Duration must be greater than 1"))
		return
	}

	member.EndDate = dateStr.AddDate(0, 0, int(member.Duration*30)).Format("02 Jan 2006")

	DB.Create(&member)

	resJSON, err := json.Marshal(member)
	if err != nil {
		panic(err)
	}
	w.Write(resJSON)

}

// MemberHandler gives the list of members present in a gym
func MemberHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	members := &[]mod.Member{}
	DB.Unscoped().Find(members)
	resJSON, err := json.Marshal(members)
	if err != nil {
		panic(err)
	}
	w.Write(resJSON)

}

func PriceGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var memberhips []mod.MemberShip
	DB.Find(&memberhips)

	json.NewEncoder(w).Encode(&memberhips)

}

// PriceUpdateHandler handle the request to update the price of type of membership of a gym
func PriceUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var memShip mod.MemberShip
	//Taking input from body
	err := json.NewDecoder(r.Body).Decode(&memShip)
	if err != nil {
		panic(err)
	}

	// DB.Where("name =?", memShip.Name).Update("price", memShip.Price)
	DB.Model(&mod.MemberShip{}).Where("name =?", memShip.Name).Updates(&memShip)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Price updated successfully"))

}

// EndMemberHandler hadles the request to end membership for a given member
func EndMemberShipHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//get the member id from the query parameters
	id := r.URL.Query().Get("id")

	//get the current time trunctaed to the nearest day
	now := time.Now().Truncate(24 * time.Hour)
	members := &[]mod.Member{}
	DB.Find(members)
	//Loop through the member list to find the given member id
	for _, member := range *members {
		//check if the id matches with the current member id
		if id == member.Member_id {

			//calculate the dduration in days
			startDate, err := time.Parse("02 Jan 2006", member.StartDate)
			if err != nil {
				log.Fatal(err)
			}
			temp := now.Sub(startDate).Hours() / 24
			duration := float64(temp)
			fmt.Println("Duration is", duration)
			// if duration < 30 {
			// 	http.Error(w, "Cannot end membership before one month", http.StatusBadRequest)
			// 	return

			// }

			oneDayMoney := (member.TotalPrice / (float64(member.Duration) * 30))
			MoneyRefund := math.Round((member.TotalPrice - (duration * oneDayMoney)) / 2)

			//Update members information
			member.Duration = temp / 30
			member.EndDate = time.Now().AddDate(0, 0, int(duration)).Format("02 Jan 2006")
			member.TotalPrice -= MoneyRefund

			DB.Where("member_id=?", id).Updates(&member)
			DB.Where("member_id=?", id).Delete(&mod.Member{})
			w.Write([]byte("User Deleted Succesfully"))

			json.NewEncoder(w).Encode(&member)
			return
		}
	}
	http.Error(w, "No member with the given ID was found", http.StatusNotFound)

}
