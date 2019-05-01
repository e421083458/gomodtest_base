package gomodtest_base

import (
	"fmt"
	"github.com/jianfengye/collection"
)

func TestFunc()  string{
	return "TextFunc"
}

func NewIntCollection()  {
	intColl := collection.NewIntCollection([]int{3,4})
	intColl.DD()
	fmt.Println("v2.0.0")
}