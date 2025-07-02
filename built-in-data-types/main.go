package main

import "fmt"

func main() {
	a, b, c, d := "happiness", 7, 45.243, '\u03B1'

	fmt.Printf("%v \t %T \n", a, a)
	fmt.Printf("%v \t %T \n", b, b)
	fmt.Printf("%v \t %T \n", c, c)
	fmt.Printf("%c \t %T \n", d, d)
}
