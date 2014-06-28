
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
	first_name string
	last_name string
}

func (p *Human) Hello() {
	fmt.Println("Hello")
}

func (p *Human) Fullname() {
	fmt.Println(p.first_name + " " + p.last_name)
}

func main() {
	a := Animal{"zzzz"}
	a.Say()
	a.Hello()
	
	aa := Ape{Animal{"apoo"}}
	aa.Say()
	aa.Hello()
	
	p := Human{Ape{Animal{"Yo"}}, "Lucy", "Sky"}
	p.Say()
	p.Hello()
	p.Fullname()
}