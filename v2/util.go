package v2

import (
	"fmt"
	"github.com/jianfengye/collection"
)

func TestFunc()  string{
	return "TextFunc"
}

func NewIntCollection(name string,sex string)  {
	intColl := collection.NewIntCollection([]int{5,6})
	intColl.DD()
	fmt.Println("name=\t",name)
	fmt.Println("sex=\t",sex)
	fmt.Println("v1.0.2")
}