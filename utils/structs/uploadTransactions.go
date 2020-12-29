package structs

type Product struct {
  ProductId string     `"json:productId"`
}

type Transaction struct {
  IdTransaction string `"json:idTransaction"`
  BuyerId string       `"json:buyerId"`
  Ip string            `"json:ip"`
  Device string        `"json:device"`
  Products []Product   `"json:products"`
}
