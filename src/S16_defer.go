package main

import "fmt"

func main() {
	defer fmt.Println("defer 1 : hihi, c'est bien moi")
	defer fmt.Println("defer 2 : non c'est moi ?")
	fmt.Println("Je suis le dernier Println, euh vraiment ?")
}
