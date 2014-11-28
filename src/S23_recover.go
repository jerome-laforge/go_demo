package main

import "fmt"

// This func print 3 lines of int array
// and sometimes its length
func printFirstLines(arr []int) {
	i := 0
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("work failed:", err)
			fmt.Println("len(arr) = ", i)
		}
		fmt.Println("-------------")
	}()

	for ; i < 3; i++ {
		fmt.Println("arr[", i, "] = ", arr[i])
	}
}

func main() {
	a4 := []int{0, 1, 3, 4}
	a2 := a4[0:2]
	a1 := a2[0:1]

	printFirstLines(a1)
	printFirstLines(a2)
	printFirstLines(a4)
	//panic("Finallement, je panique")
}
