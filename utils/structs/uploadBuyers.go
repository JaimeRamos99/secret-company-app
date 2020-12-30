package structs

type Buyer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewBuyer(id string, name string, age int) *Buyer {
	return &Buyer{Id: id, Name: name, Age: age}
}
