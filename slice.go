package main

import "fmt"

func main() {

	var as []int = []int{1, 2, 3}

	as = append(as, 4)

	fmt.Println(len(as))

	as = make([]int, 1, 4)
	fmt.Println(len(as), cap(as))

	x := [3]string{"Лайка", "Белка", "Стрелка"}
	s := x[:] // a slice referencing the storage of x

	s = append([]string{}, x[:]...)

	for i, e := range s {
		fmt.Printf("s[%d]=%s\n", i, e)
	}
}
