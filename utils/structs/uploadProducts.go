package structs

type Product struct {
	Uid         string `json:"uid,omitempty"`
	ProductID   string `json:"productId"`
	ProductName string `json:"productName"`
	Price       int    `json:"price"`
}

func NewProduct(productID string, productName string, price int) *Product {
	return &Product{Uid: "_:" + productID, ProductID: productID, ProductName: productName, Price: price}
}
