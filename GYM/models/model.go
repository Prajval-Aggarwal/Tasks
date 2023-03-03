package mod

type Members struct {
	ID           int     `json:"id"`
	StartDate    string  `json:"start_date"`
	Name         string  `json:"name"`
	MemberShip   string  `json:"member_ship"`
	TotalPrice   float64 `json:"total_price"`
	EndDate      string  `json:"end_date"`
	MonthlyPrice float64 `json:"monthly_price"`
	Duration     int     `json:"duration"`
	IsDelete     bool    `json:"is_delete"`
	// MoneySubmitted float64 `json:"money_submitted"`
	// Balance        float64 `json:"balance"`
}

type Exit struct {
	Duration float64 `json:"duration"`
	Refund   float64 `json:"refund"`
}
