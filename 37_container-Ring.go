package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5) // tentukan jumlah data ring nya

	data.Value = "bambang"; // untuk memasukkan data ke ring

	for i := 0; i < data.Len(); i++ {
		data.Value = "herman" + strconv.FormatInt(int64(i), 10);
		data = data.Next()
	}

	data.Do(func(value interface{}){ // gunakan function Do, untuk mengiterasi data (menampilkan data nya)
		fmt.Println(value)
	})
}