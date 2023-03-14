package mod

import "gorm.io/gorm"

//import "github.com/golang-jwt/jwt/v4"

type Member struct {
	Member_id       string `json:"id" gorm:"default:uuid_generate_v4()"`
	StartDate       string `json:"start_date"`
	Name            string `json:"name"`
	MemberShip_Name string `json:"member_ship_name"`
	// MemberShip      MemberShip     `gorm:"foreignKey:MemberShip_Name;references:Name"`
	TotalPrice   float64        `json:"total_price"`
	EndDate      string         `json:"end_date"`
	MonthlyPrice float64        `json:"monthly_price"`
	Duration     float64        `json:"duration"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type MemberShip struct {
	Name  string  `json:"name" gorm:"unique"`
	Price float64 `json:"price"`
}

type User struct {
	User_Id   string `json:"user_id" gorm:"default:uuid_generate_v4();uniqueIndex"` //PK
	User_Name string `json:"name"`
	Gender    string `json:"gender"`
}

type Payment struct {
	Payment_Id   string  `json:"payment_id" gorm:"default:uuid_generate_v4();unique"` //PK
	User_Id      string  `json:"user_id"`                                             //FK
	User         User    `gorm:"references:User_Id"`
	Amount       float64 `json:"amount"`
	Payment_Type string  `json:"payment_type"`
}

type Subscription struct {
	Payment_Id string `json:"payment_id"` //FK
	//Payment    Payment `gorm:"references:Payment_Id"`
	User_Id string `json:"user_id"` //Fk
	User    User   `gorm:"references:User_Id"`
	Emp_Id  string `json:"emp_id"` //FK
	//Emp        GymEmp  `gorm:"references:Emp_Id"`
	Subs_Name string         `json:"subs_name"` //FK
	StartDate string         `json:"start_date"`
	EndDate   string         `json:"end_date"`
	Duration  float64        `json:"duration"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type SubsType struct {
	Subs_Name string  `json:"name" gorm:"unique"`
	Price     float64 `json:"price"`
}
type GymEmp struct {
	Emp_Id   string `json:"emp_id" gorm:"default:uuid_generate_v4();unique"` //PK
	Emp_name string `json:"emp_name"`
	Gender   string `json:"gender"`
	Role     string `json:"role"`
}

type Equipment struct {
	Model_No   string `json:"model_no" gorm:"uuid_generate_v4();unique"` //PK
	Equip_Name string `json:"equip_name"`
	Quantity   int64  `json:"quantity"`
}
