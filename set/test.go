package main

import "fmt"

func main() {
	set := NewSet()
	array := []int{1, 4, 5, 9, 1, 9, 0, 2}
	for _, a := range array {
		set.Add(a)
	}
	fmt.Println(set.Len())
	fmt.Println(set.Contains(2))
	fmt.Println(set.Contains(3))
	set.Remove(9)
	fmt.Println(set.Len())
}
