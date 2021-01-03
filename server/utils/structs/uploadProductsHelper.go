package structs

type ProductIdResp struct {
	Uid       string `json:"uid"`
	ProductId string `json:"productId"`
}
type Products struct {
	Products []ProductIdResp `json:"products"`
}
