package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data);

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i);

		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface(); // objectnya.field nya keberapa.conversi ke interface
			if value == "true" {
				return false
			}
		}
	}

	return true;
}

func main() {
	sample := Sample{Name: "Akbar"}

	sampleType := reflect.TypeOf(sample);

	fmt.Println(sampleType.NumField()) // mengetahui ada berapa jumlah fieldnya
	fmt.Println(sampleType.Field(0).Name) // mengetahui nama field 0

	fmt.Println(sampleType.Field(0).Tag.Get("required")) // mengambil tag
	fmt.Println(sampleType.Field(0).Tag.Get("max")) // mengambil tag


	fmt.Println(isValid(sample))

}