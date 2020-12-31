package structs

//Struct for the query response
type Resp struct {
	Date string `json:"date"`
}

type RespArray struct {
	Query []Resp `json:"query"`
}

func NewResp(date string) *Resp {
	return &Resp{Date: date}
}
