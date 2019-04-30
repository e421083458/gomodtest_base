package main

import "github.com/jianfengye/collection"

func main()  {
	intColl := collection.NewIntCollection([]int{1,2})
	intColl2 := intColl.NewEmpty()
	intColl2.DD()
	/*
	IntCollection(0):{
	}
	*/
}
