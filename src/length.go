package main

import "fmt"


func main() {
	test := map[string]interface{}{
		"1": 2,
		"2": 3,
		"a": 5}
		
	fmt.Println(len(test))
}