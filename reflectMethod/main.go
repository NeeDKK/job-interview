package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Name string `field:"namePeople"`
}

func (p *People) Eat() {
	fmt.Println("people eat")
}

func main() {
	of := reflect.ValueOf(&People{})
	of.MethodByName("Eat").Call([]reflect.Value{})
	name, _ := reflect.TypeOf(People{}).FieldByName("Name")
	get := name.Tag.Get("field")
	fmt.Println(get)
}
