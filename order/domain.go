package order

// To accept from body
type Item struct {
	Product_Id int
	Quantity   int
}

type Order struct {
	Item []Item
}

// To display orders
type Order1 struct {
	ID           int
	Amount       int
	Disc_perc    int
	Final_amnt   int
	Disp_date    string `json:"dispatch_date,omitempty"`
	Order_status string
}
type ListResponse struct {
	Orders []Order1 `json:"order1"`
}

type FindByIdResponse struct {
	Order Order1 `json:"order1"`
}
