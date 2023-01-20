package main

import "fmt"

func addReturn() *int {
	v := 5
	return &v
}

func main() {
	i := 1
	p := &i

	*p = *p + 2

	fmt.Println(i, p)
	fmt.Println(addReturn)
}
