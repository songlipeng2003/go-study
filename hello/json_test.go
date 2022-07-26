package hello

import "testing"

func TestJson(t *testing.T) {
	user := User{"Hello", "Thinking", []string{"Go", "Python", "JavaScript"}}
	u, _ := ToJson(user)
	s := string(u)

	if s != `{"Name":"Hello","Website":"Thinking","Skills":["Go","Python","JavaScript"]}` {
		t.Errorf("json is error: %s, ", u)
	}

	user2, _ := ToObject(s)

	if user2.Name != user.Name {
		t.Errorf("json is error: %s, ", u)
	}
}
