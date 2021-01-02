package structs

type IP struct {
	Ip string `json:"ip,omitempty"`
}

type FirstStage struct {
	MadeBy []IP `json:"~madeBy,omitempty"`
}

type UserSameIp struct {
	UserId   string `json:"userId,omitempty"`
	UserName string `json:"userName,omitempty"`
}

type SecondStage struct {
	TransactionId string       `json:"transactionId,omitempty"`
	Ip            string       `json:"ip,omitempty"`
	MadeBy        []UserSameIp `json:"madeBy,omitempty"`
}

type DataSameIps struct {
	First    []FirstStage  `json:"first_stage,omitempty"`
	StageTwo []SecondStage `json:"second_stage,omitempty"`
}
