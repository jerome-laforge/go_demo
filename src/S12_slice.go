package main

import (
	"fmt"
)

func main() {
	s0 := []byte{0, 1, 2, 3, 4}
	s1 := s0[2:4]

	fmt.Println("s0 [", &s0[0], "] =", s0)
	fmt.Println("s1 [", &s1[0], "] =", s1)
}
