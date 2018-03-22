package main

import "fmt"

type Shape interface {
	area() int      // Méthode "area" renvoyant un "int"
	perimeter() int // Méthode "perimeter" renvoyant un "int"
}

type Rectangle struct {
	width, height int // Déclaration des champs "width" et "height" de type "int"
}

// Déclaration de Rectangle::area(), renvoyant un "int"
func (r Rectangle) area() int {
	return r.width * r.height
}

// Déclaration de Rectangle::perimeter(), renvoyant un "int"
func (r Rectangle) perimeter() int {
	return r.width*2 + r.height*2
}

func foo(s Shape) {
	fmt.Println(s.area())      // Affichage de son aire
	fmt.Println(s.perimeter()) // Affichage de son périmètre
}

func main() {
	r := Rectangle{10, 20} // Déclaration d'un rectangle
	foo(r)                 // Appel de "foo" avec "r"
	// ce dernier étant de type "Rectangle"
	// et "Rectangle" héritant de "Shape"
}
