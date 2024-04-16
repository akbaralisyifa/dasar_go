package main

import "fmt"

type Man struct {
	name string
}

func (man *Man) GetGender() {
	man.name = "Mr " + man.name
}

func main() {
	eko := Man{"Eko"}
	eko.GetGender()

	fmt.Println(eko)

}