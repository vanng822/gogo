package main

import (
	"fmt"
)

type Car struct {
	engine         string
	numberOfWheels int
}

type Boat struct {
	engine       string
	drivingForce string
}

type AmphibiousVehicle struct {
	Car
	Boat
}

func main() {
	av := AmphibiousVehicle{Car{engine: "v8", numberOfWheels: 4}, Boat{engine: "disel", drivingForce: "pppp"}}
	fmt.Println(av)
}
