package gomodtest_base

import (
	"fmt"
	"github.com/jianfengye/collection"
)

func TestFunc()  string{
	return "TextFunc"
}

func NewIntCollection()  {
	intColl := collection.NewIntCollection([]int{1,2})
	intColl.DD()
	fmt.Println("v1.0.0")
}