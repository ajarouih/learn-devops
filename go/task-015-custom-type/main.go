package main

import "fmt"

var a int

type testtype int

var b testtype

func main() {
	a = 23
	b = 34
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("\n%T", a)
	fmt.Printf("\n%T", b)

}