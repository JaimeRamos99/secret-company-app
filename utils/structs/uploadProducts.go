package structs

type Product struct {
	Uid         string `json:"uid,omitempty"`
	ProductID   string `json:"productId,omitempty"`
	ProductName string `json:"productName,omitempty"`
	Price       int    `json:"price,omitempty"`
}

func NewProduct(productID string, productName string, price int) *Product {
	return &Product{Uid: "_:" + productID, ProductID: productID, ProductName: productName, Price: price}
}
