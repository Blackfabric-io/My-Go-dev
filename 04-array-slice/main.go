package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "John"
	names[1] = "Doe"
	names[2] = "Frank"
	names[3] = "Jack"

	fmt.Println(names) // Output: [John Doe Frank Jack]
}
