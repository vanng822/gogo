package main

import (
	"time"
	"strconv"
	"fmt"
)

func main() {
	start := time.Now()
	mapString := make(map[string]int, 300)
	for i := 0; i < 300; i++ {
		mapString[strconv.Itoa(i)] = i
	}
	
	for j := 0; j < 1000; j++ {
		_ = mapString["100"]
	}
	fmt.Println(time.Now().Sub(start))

	starti := time.Now()
	mapInt := make(map[int]int, 300)
	for i := 0; i < 300; i++ {
		mapInt[i] = i
	}
	for j := 0; j < 1000; j++ {
		i, _ := strconv.Atoi("100")
		_ = mapInt[i]
	}
	fmt.Println(time.Now().Sub(starti))	
}

