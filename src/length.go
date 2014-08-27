package main

import "fmt"

type D map[string]interface{}

type T struct {
	t []D
}

func main() {
	test := map[string]interface{}{
		"1": 2,
		"2": 3,
		"a": 5}
		
	fmt.Println(len(test))
	
	t := new(T)
	
	fmt.Println(len(t.t))
}