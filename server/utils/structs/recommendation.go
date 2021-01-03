package structs

type Counter struct {
	Score       int    `json:"score,omitempty"`
	ProductName string `json:"productName,omitempty"`
}

type Includes struct {
	Includes []Counter `json:"includes,omitempty"`
}

type Result struct {
	Result []Includes `json:"result,omitempty"`
}
