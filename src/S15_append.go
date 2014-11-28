package main

import "fmt"

func main() {
	s0 := []byte{0, 1}
	fmt.Println("length s0 = ", len(s0))
	fmt.Println("capacity s0 = ", cap(s0))
	fmt.Println("s0 [", &s0[0], "] = ", s0)

	fmt.Println("")
	s0 = append(s0, 2, 3)
	fmt.Println("length s0 = ", len(s0))
	fmt.Println("capacity s0 = ", cap(s0))
	fmt.Println("s0 [", &s0[0], "] = ", s0)

	fmt.Println("")
	s1 := make([]byte, 1, 5)
	s1[0] = 255
	fmt.Println("length s1 = ", len(s1))
	fmt.Println("capacity s1 = ", cap(s1))
	fmt.Println("s1 [", &s1[0], "] = ", s1) // essayons avec un len = 0?

	fmt.Println("")
	s1 = append(s1, s0...)
	fmt.Println("length s1 = ", len(s1))
	fmt.Println("capacity s1 = ", cap(s1))
	fmt.Println("s1 [", &s1[0], "] = ", s1)
//
//	fmt.Println("")
//	s1 = append(s1, s0[0])
//	fmt.Println("length s1 = ", len(s1))
//	fmt.Println("capacity s1 = ", cap(s1))
//	fmt.Println("s1 [", &s1[0], "] = ", s1)
}
