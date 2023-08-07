package order

type Item struct {
	Product_Id int
	Quantity   int
}

type Order struct {
	Item []Item
}

// type ListResponse struct {
// 	Orders []Order
// }
