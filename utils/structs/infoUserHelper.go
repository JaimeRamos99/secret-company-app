package structs

type TopProduct struct {
	ProductName string `json:"name,omitempty"`
	Score       int    `json:"score,omitempty"`
}

func NewTopProduct(productName string, score int) *TopProduct {
	return &TopProduct{
		ProductName: productName,
		Score:       score,
	}
}
