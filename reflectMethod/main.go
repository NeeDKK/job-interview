package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Name string
}

func (p *People) Eat() {
	fmt.Println("people eat")
}

func main() {
	of := reflect.ValueOf(&People{})
	of.MethodByName("Eat").Call([]reflect.Value{})
}
