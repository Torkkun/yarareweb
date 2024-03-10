package domain

type Data struct {
	ProductID    string `json:"productId"`
	ProductName  string `json:"productName"`
	UserName     string `json:"userName"`
	PurchaseDate string `json:"purchaseDate"`
}

type DataResponse []*Data
