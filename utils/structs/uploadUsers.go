package structs

type Buyer struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type User struct {
	Uid    string `json:"uid,omitempty"`
	UserId string `json:"userId,omitempty"`
	Name   string `json:"userName,omitempty"`
	Age    int    `json:"userAge,omitempty"`
}

type Users struct {
	Users []User
}

func NewBuyer(id string, name string, age int) *Buyer {
	return &Buyer{
		Id:   id,
		Name: name,
		Age:  age,
	}
}

func NewUser(userId string, name string, age int) *User {
	return &User{
		Uid:    "_:" + userId,
		UserId: userId,
		Name:   name,
		Age:    age,
	}
}
