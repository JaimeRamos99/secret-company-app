package structs

type Date struct {
	Date  string   `json:"date,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}
