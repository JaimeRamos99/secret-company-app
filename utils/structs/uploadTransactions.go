package structs

type ProductId struct {
  ProductId string     `"json:productId"`
}

type Transaction struct {
  IdTransaction string `"json:idTransaction"`
  BuyerId string       `"json:buyerId"`
  Ip string            `"json:ip"`
  Device string        `"json:device"`
  Products []ProductId   `"json:products"`
}

func NewProductId(productId string) *ProductId {
	return &ProductId{ ProductId: productId}
}

func NewTransaction(idTransaction string, buyerId string, ip string, device string, products []ProductId) *Transaction {
	return &Transaction{IdTransaction: idTransaction, BuyerId: buyerId, Ip: ip, Device: device, Products: products}
}
