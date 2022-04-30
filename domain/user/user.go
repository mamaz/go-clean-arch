package user

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

func (usr *User) Hash() string {
	hasher := sha1.New()

	str := fmt.Sprintf("%s:%s:%s", usr.ID, usr.Name, usr.City)
	hasher.Write([]byte(str))

	return hex.EncodeToString(hasher.Sum(nil))
}
