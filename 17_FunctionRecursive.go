package main

import "fmt"

// Menggunakan Looping
func factorialLoop(value int) int {
	result := 1;
	for i:= value; i > 0; i-- {
		result *= i;  // result = result * i
	}

	return result;
}

// Menggunakan recursive
func factorialRecursive(value int) int {
	if value == 1{
		return 1;
	}else{
		return value * factorialRecursive(value - 1);
	}
}

func main(){
	loop := factorialLoop(5);
	fmt.Println(loop);
	
	fmt.Println("menggunakan recursive");
	recursive := factorialRecursive(5);
	fmt.Println(recursive)

}