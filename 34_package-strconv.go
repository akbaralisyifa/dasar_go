package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// conversi dari string boolean ke bolearn
	boolean, err := strconv.ParseBool(strings.ToLower("True")) // untuk mengkonversi dari string ke boolean

	if err == nil {
		fmt.Println(boolean)
	}else{
		fmt.Println(err.Error())
	}

	// conversi dari string ke int
	numver, errNumber := strconv.ParseInt("100", 10, 64);
	if errNumber == nil {
		fmt.Println(numver)
	}else{
		fmt.Println(errNumber.Error())
	}

}