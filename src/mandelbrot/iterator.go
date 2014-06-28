package main

import "fmt"

const BAILOUT float64 = 16
const MAX_ITERATIONS int = 1000

func mandelbrot_main() {
	var y int
	var x int
	var i int
	for y = -39; y < 39; y++ {
		fmt.Printf("\n")
		for x = -39; x < 39; x++ {
			i = mandelbrot(float64(x)/40.0, float64(y)/40.0)
			if i == 0 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
	}
}

func mandelbrot(x float64, y float64) int {
	cr := y - 0.5
	ci := x
	zi := 0.0
	zr := 0.0
	i := 0

	for {
		i += 1
		temp := zr * zi
		zr2 := zr * zr
		zi2 := zi * zi
		zr = zr2 - zi2 + cr
		zi = temp + temp + ci

		if zi2+zr2 > BAILOUT {
			break
		}

		if i > MAX_ITERATIONS {
			i = 0
			break
		}
	}

	return i
}

func main() {
	mandelbrot_main()
	fmt.Printf("\n")
}
