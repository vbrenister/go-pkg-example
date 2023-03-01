package users

import "fmt"

type User struct {
	Name string
	Age int
}


func (u User) String() string {
	return fmt.Sprintf("User %v with age %d\n", u.Name, u.Age)
}

func (u User) GetGreetings() string {
	return fmt.Sprintf("%v says greetings\n", u.Name)
}