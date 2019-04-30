package gomodtest_base

import "github.com/jianfengye/collection"

func TestFunc()  string{
	return "TextFunc"
}

func NewIntCollection()  {
	intColl := collection.NewIntCollection([]int{1,2})
	intColl2 := intColl.NewEmpty()
	intColl2.DD()
}