package product

type Product struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	Availability int32  `json:"availability"`
	Price        int32  `json:"price"`
	Category     string `json:"category"`
}

type ListResponse struct {
	Products []Product `json:"product"`
}
