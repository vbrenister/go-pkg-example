package greeter

import "fmt"

type Greeter interface {
	GetGreetings() string
}


func Greet(greeter Greeter) {
	fmt.Println(greeter.GetGreetings())
}