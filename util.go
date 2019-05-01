package gomodtest_base

import (
	"fmt"
	"github.com/jianfengye/collection"
)

func TestFunc()  string{
	return "TextFunc"
}

func NewIntCollection(name string)  {
	intColl := collection.NewIntCollection([]int{3,4})
	intColl.DD()
	fmt.Println("name=\t",name)
	fmt.Println("v1.0.1")
}