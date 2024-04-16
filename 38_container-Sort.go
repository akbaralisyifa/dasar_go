package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int { // func len ada lah angka dari element di dalam data slice
	return len(value)
}

func (value UserSlice) Less(i, j int) bool { // mengurut kan data dari yang terkecil
	return value[i].Age < value[j].Age
}

func (value UserSlice) swap(i, j int) { // untuk mengurutkan secara terbalik
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"herman", 22},
		{"bambang", 20},
		{"budi", 30},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)

}