package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Day())
	fmt.Println(now.Month())
	fmt.Println(now.Year())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC);
	fmt.Println(utc);

	// untuk parse dari string ke time
	layout := time.RFC1123;
	parse, _ := time.Parse(layout, "01/02/2006 15:04:05");
	fmt.Println(parse)
}