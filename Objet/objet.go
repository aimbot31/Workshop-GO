/*
** EPITECH PROJECT, 2017
** objet.go
** File description:
** florian.davasse@epitech.eu
 */

package main

import "fmt"

type Person struct {
	Name string
	Address
}

type Address struct {
	Number string
	Street string
	City   string
	State  string
	Zip    string
}

type Citizen struct {
	Country string
	Person
}

func (c *Citizen) Nationality() {
	fmt.Println(c.Name, "is a citizen of", c.Country)
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func main() {
	c := Citizen{}
	c.Name = "Steve"
	c.Country = "America"
	c.Talk()        // <- Notice both are accessible
	c.Person.Talk() // <- Notice both are accessible
	c.Nationality()
}
