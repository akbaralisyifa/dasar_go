package main

import "fmt"

// Funciton
func myWorld() {
	fmt.Println("My World Is You")
}

// Function with parameter
func myHello (name string){
	fmt.Println("My World Is You : ", name)
}

	// Function Return Value
	func myReturn(name string) string { // string ini di gunakan untuk tipe return value
		if name == "" {
			return "My World Is You"
		}else{
			return "Hello " + name;
		}
	}

	// function multiple return value
	func renturnMuliple()(string, string){
		return "bambang", "herman";
	}

	// Function return named value
	func fullName()(firstName string, lastName string){
		// initial value (default value)
		firstName = "bambang"
		lastName = "herman"

		// return firstName, lastName
		// just return 
		return
	}

	// Variadic Function
	func sumAll(numbers ...int)int {
		total := 0;
		for _, value := range numbers {
			total += value;
		}

		return total;
	}

	// function Value
	func sayGoodBye(name string) string {
		return "Good Bye " + name;
	}

func main() {
	myWorld();
	myHello("barre")
	fmt.Println(myReturn(("Kmu")))

	// multiple return value
	firstName, lastName := renturnMuliple();
	fmt.Println(firstName, lastName);
	// or
	nameFirst, _ := renturnMuliple();
	fmt.Println(nameFirst);


	// return named 
	nameOne, nameTwo := fullName();
	fmt.Println(nameOne, nameTwo);

	// Variadic Function 
	total := sumAll(11,23, 23);
	fmt.Println(total)

	slice := []int{112,11,22};
	result := sumAll(slice...);
	fmt.Println(result)


	// Function Value
	getGoodBye := sayGoodBye;
	result := getGoodBye("bambang");
	fmt.Println(result)

}