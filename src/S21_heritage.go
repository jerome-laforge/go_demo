package main

import "fmt"

type Hobby interface {
	myType() string
	myStereotype() string
}

type Human struct {
}

func (h Human) myType() string {
	return "Je suis un humain"
}

func (_ Human) myStereotype() string {
	return "je suis conceptuel, et je n'ai pas de passe-temps."
}

type Man struct {
	Human //anonymous class to inherit Human behavior
}

func (_ Man) myStereotype() string {
	return "j'aime les voitures."
}

type Woman struct {
	Human //anonymous class to inherit Human behavior
}

func (_ Woman) myStereotype() string {
	return "j'aime le shopping."
}

type Dog struct {
	//does not inherit any other type
}

func (_ Dog) myType() string {
	return "Je suis un chien"
}

func (_ Dog) myStereotype() string {
	return "je vais chercher des bâtons."
}

func main() {
	h := new(Human)
	m := new(Man)
	w := new(Woman)
	d := new(Dog)

	//an array of hobby instances - we don’t need to know whether human or dog
	hobbyArr := [...]Hobby{h, m, w, d} //array of 3 Humans and 1 dog.
	for k, v := range hobbyArr {
		fmt.Println(v.myType())
		fmt.Println("Mes passe-temps?  Bien,", hobbyArr[k].myStereotype()) //appears as Hobby type, but behavior changes depending on actual instance

	}
}
