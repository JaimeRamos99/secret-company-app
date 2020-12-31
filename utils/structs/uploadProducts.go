package structs

type Product struct {
	ProductID   string `json:"productId"`
	ProductName string `json:"productName"`
	Price       int    `json:"price"`
}

func NewProduct(productID string, productName string, price int) *Product {
	return &Product{ProductID: productID, ProductName: productName, Price: price}
}
