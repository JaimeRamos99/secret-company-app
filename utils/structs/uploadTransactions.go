package structs

type ProductId struct {
	ProductId string `json:"productId"`
}

type Transaction struct {
	IdTransaction string      `json:"idTransaction"`
	BuyerId       string      `json:"buyerId"`
	Ip            string      `json:"ip"`
	Device        string      `json:"device"`
	Products      []ProductId `json:"products"`
}

//buyer for transaction json
type MadeBy struct {
	Uid string `json:"uid"`
}

type ProductUid struct {
	Uid string `json:"uid"`
}

type TransactionMutation struct {
	MadeBy        MadeBy       `json:"madeBy"`
	IdTransaction string       `json:"transactionId"`
	Ip            string       `json:"ip"`
	Device        string       `json:"device"`
	Includes      []ProductUid `json:"includes"`
}

func NewMadeBy(userUid string) *MadeBy {
	return &MadeBy{Uid: userUid}
}

func NewProductUId(productUid string) *ProductUid {
	return &ProductUid{Uid: productUid}
}

func NewProductId(productId string) *ProductId {
	return &ProductId{ProductId: productId}
}

func NewTransaction(idTransaction string, buyerId string, ip string, device string, products []ProductId) *Transaction {
	return &Transaction{
		IdTransaction: idTransaction,
		BuyerId:       buyerId,
		Ip:            ip,
		Device:        device,
		Products:      products,
	}
}

func NewTransactionMutation(madeBy MadeBy, idTransaction string, ip string, device string, products []ProductUid) *TransactionMutation {
	return &TransactionMutation{
		MadeBy:        madeBy,
		IdTransaction: idTransaction,
		Ip:            ip,
		Device:        device,
		Includes:      products,
	}
}
