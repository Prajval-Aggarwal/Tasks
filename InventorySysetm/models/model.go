package mod

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

type Cart struct {
	Prod  []Product
	Total float64
}
