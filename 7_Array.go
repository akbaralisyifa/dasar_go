package main

import "fmt";

func main(){
	var names [4]string;

	names[0] = "Ali";
	names[1] = "Budi";
	names[2] = "Caca";

	fmt.Println(names[0], names[1], names[2]);

	// or

	var umur = []int{1,3,4};
	
	fmt.Println(umur)
	fmt.Println(umur[0])

	fmt.Println(umur[1] = 2)

	// length
	fmt.Println(len(umur))

}