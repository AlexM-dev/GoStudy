package main

import "fmt"

func main() {
	var s1, s2 int
	fmt.Println("Enter first array size:")
	fmt.Scanln(&s1)
	fmt.Println("Enter second array size:")
	fmt.Scanln(&s2)
	m1 := make([]string, s1)
	m2 := make([]string, s2)

	fmt.Println("Enter first array:")
	for i := 0; i < s1; i++ {
		fmt.Scanln(&m1[i])
	}
	fmt.Println("Enter second array")
	for i := 0; i < s2; i++ {
		fmt.Scanln(&m2[i])
	}
	m := make(map[string]int)
	for _, i := range m1 {
		for _, j := range m2 {
			if i == j {
				m[i]++
			}
		}
	}
	fmt.Println("Equal elements:")
	for key, _ := range m {
		fmt.Print(key, " ")
	}
}
