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
