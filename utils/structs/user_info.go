package structs

type UserId struct {
	Id string `json:"id,omitempty"`
}

type ProductsResp struct {
	ProductPrice string `json:"productPrice,omitempty"`
	ProductName  string `json:"productName,omitempty"`
	ProductId    string `json:"productId,omitempty"`
}

type TransactionsResp struct {
	Ip            string         `json:"ip,omitempty"`
	TransactionId string         `json:"transactionId,omitempty"`
	Products      []ProductsResp `json:"includes,omitempty"`
}

type Info struct {
	UserId       string             `json:"userId,omitempty"`
	UserAge      string             `json:"userAge,omitempty"`
	UserName     string             `json:"userName,omitempty"`
	Transactions []TransactionsResp `json:"transactions,omitempty"`
}
type InfoArray struct {
	InfoArray []Info `json:"info,omitempty"`
}
