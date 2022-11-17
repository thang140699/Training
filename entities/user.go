package entities

import (
	"fmt"
)

type User struct {
	Id     string `json:"Id" bson:"Id"`
	Time   int64  `json:"Time" bson:"Time"`
	Domain string `json:"Domain" bson:"Domain"`
}

func (user User) ToString() string {
	return fmt.Sprintf("Id: %s\nId: %s\nDomain: %s\n", user.Id, user.Domain)
}
