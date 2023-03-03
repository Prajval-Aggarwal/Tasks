package mod

//import "github.com/golang-jwt/jwt/v4"

type Members struct {
	ID           int     `json:"id"`
	StartDate    string  `json:"start_date"`
	Name         string  `json:"name"`
	MemberShip   string  `json:"member_ship"`
	TotalPrice   float64 `json:"total_price"`
	EndDate      string  `json:"end_date"`
	MonthlyPrice float64 `json:"monthly_price"`
	Duration     float64 `json:"duration"`
	IsDelete     bool    `json:"is_delete"`
}
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Exit struct {
	Duration float64 `json:"duration"`
	Refund   float64 `json:"refund"`
}

// type Claims struct {
// 	Username string `json:"username"`
// 	jwt.RegisteredClaims
// }
