package hello

import "encoding/json"

type User struct {
	Name    string
	Website string
	Skills  []string
}

func ToJson(user User) ([]byte, error) {
	u, err := json.Marshal(user)
	return u, err
}

func ToObject(s string) (User, error) {
	var user User
	err := json.Unmarshal([]byte(s), &user)
	return user, err
}
