package main

import "fmt"

func main(){
	type noKTP string;
	type meried bool;

	var noKTPAli noKTP = "1234567890123456";
	var isMerrid meried = true;

	fmt.Println(noKTPAli, isMerrid)
}