package main

import "fmt"

type str string

func (text str) log() {
	fmt.Println(text)
}

func main() {
	var name str = "Ray" // if you want to use your custom type you have to set it explicitly

	name.log()
}
