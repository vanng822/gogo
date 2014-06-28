
package main

import "fmt"

type Animal struct {
	sound string
}

func (a *Animal) Say() {
	fmt.Println(a.sound)
}

func (a *Animal) Hello() {
	fmt.Println("Uh Uh")
}

type Ape struct {
	Animal
}

func (p *Ape) Hello() {
	fmt.Println("Whaaa")
}

type Human struct {
	Ape
}

func (p *Human) Hello() {
	fmt.Println("Hello")
}

func main() {
	a := Animal{"zzzz"}
	a.Say()
	a.Hello()
	
	aa := Ape{Animal{"apoo"}}
	aa.Say()
	aa.Hello()
	
	p := Human{Ape{Animal{"Yo"}}}
	p.Say()
	p.Hello()
}