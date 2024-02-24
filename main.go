package main

import (
	"fmt"

	"github.com/amitramachandran/dsa/linkedList"
)

func printLL(temp *linkedList.Node) {
	fmt.Printf("%v --> ", temp.GetData())
	for temp.GetNext() != nil {
		temp = temp.GetNext()
		fmt.Printf("%v --> ", temp.GetData())
	}
	fmt.Print("\n")

}

func main() {
	a := linkedList.NewNode("1")
	b := linkedList.NewNode("2")
	c := linkedList.NewNode("3")
	d := linkedList.NewNode("0")
	e := linkedList.NewNode("2.5")
	a.SetHead()
	a.Connect(b)
	b.Connect(c)
	d.Connect(a)
	b.Connect(e)
	d.Delete(d)
	printLL(d)
	printLL(a)
	printLL(b)

}
