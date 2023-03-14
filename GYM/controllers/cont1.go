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

var DB1 *gorm.DB

func init() {
	fmt.Println("Connecting to database...")
	dsn := "host=localhost port=5432 user=postgres password=Test@123 dbname=gym1 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in connecting to database:", err)
	}
	db.AutoMigrate(&mod.Subscription{}, &mod.Payment{}, &mod.SubsType{}, &mod.User{}, &mod.GymEmp{}, &mod.Equipment{})

	DB1 = db
	fmt.Println("Successfully connected to database")
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user mod.User
	json.NewDecoder(r.Body).Decode(&user)
	DB1.Create(&user)
	json.NewEncoder(w).Encode(&user)
}

func MakepaymentHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("Id is :", id)
	w.Header().Set("Content-Type", "application/json")
	var payment mod.Payment
	var sub mod.Subscription
	json.NewDecoder(r.Body).Decode(&payment)

	DB1.Where("user_id=?", id).First(&sub)

	var memShip mod.SubsType
	DB1.Where("subs_name=?", sub.Subs_Name).First(&memShip)
	//fmt.Println("svjhfvvsjd:", memShip)

	payment.Amount = float64(sub.Duration) * memShip.Price
	payment.User_Id = id

	fmt.Println("payment.User.User_Id", payment.User.User_Id)

	DB1.Create(&payment)
	fmt.Println("Payment id is:", payment.Payment_Id)
	sub.Payment_Id = payment.Payment_Id
	DB1.Where("user_id=?", id).Updates(&sub)
	json.NewEncoder(w).Encode(&payment)

}

func CreateSubsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var sub mod.Subscription
	json.NewDecoder(r.Body).Decode(&sub)
	id := r.URL.Query().Get("id")
	fmt.Println("id is", id)
	dateStr := time.Now().Truncate(time.Hour)
	sub.StartDate = dateStr.Format("02 Jan 2006")

	sub.EndDate = dateStr.AddDate(0, 0, int(sub.Duration*30)).Format("02 Jan 2006")

	sub.User_Id = id
	DB1.Create(&sub)
	json.NewEncoder(w).Encode(&sub)

}

func CreateEmphandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp mod.GymEmp
	json.NewDecoder(r.Body).Decode(&emp)

	DB1.Create(&emp)

	json.NewEncoder(w).Encode(emp)
}

func AddEquipHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var equipment mod.Equipment
	json.NewDecoder(r.Body).Decode(&equipment)
	DB1.Create(&equipment)
	json.NewEncoder(w).Encode(&equipment)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var show []mod.Subscription

	DB1.Joins("User").Find(&show)

	json.NewEncoder(w).Encode(&show)

}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []mod.GymEmp

	DB1.Find(&employees)

	json.NewEncoder(w).Encode(&employees)

}

func GetEquipList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var equipments []mod.Equipment

	DB1.Find(&equipments)

	json.NewEncoder(w).Encode(&equipments)

}
func GetPrices(w http.ResponseWriter, r *http.Response) {
	w.Header().Set("Content-Type", "application/json")
	var memberhips []mod.SubsType
	DB1.Find(&memberhips)

	json.NewEncoder(w).Encode(&memberhips)
}

func PriceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var memShip mod.SubsType
	//Taking input from body
	err := json.NewDecoder(r.Body).Decode(&memShip)
	if err != nil {
		panic(err)
	}

	// DB.Where("name =?", memShip.Name).Update("price", memShip.Price)
	DB.Model(&mod.SubsType{}).Where("name =?", memShip.Subs_Name).Updates(&memShip)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Price updated successfully"))

}

func EndSubscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	now := time.Now().Truncate(24 * time.Hour)
	var subs mod.Subscription
	var payment mod.Payment
	DB1.Where("user_id=?", id).First(&subs)
	if subs.Payment_Id == "" {
		fmt.Println("Payment not done")
		DB1.Where("user_id=?", id).Delete(&subs)
		return
	}
	DB1.Where("payment_id=?", subs.Payment_Id).First(&payment)
	startDate, err := time.Parse("02 Jan 2006", subs.StartDate)
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
	oneDayMoney := (payment.Amount / (float64(subs.Duration) * 30))
	MoneyRefund := math.Round((payment.Amount - (duration * oneDayMoney)) / 2)
	subs.Duration = duration / 30
	subs.EndDate = time.Now().AddDate(0, 0, int(duration)).Format("02 Jan 2006")
	payment.Amount -= MoneyRefund
	DB1.Where("user_id=?", id).Updates(&payment)
	DB1.Where("user_id=?", id).Updates(&subs)
	DB1.Where("user_id=?", id).Delete(&subs)
	w.Write([]byte("Deleted user sucessfully.."))

}
