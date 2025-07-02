package main

import "fmt"

const (
	_ = iota
	a
	b
	c
	d
	e
	f
)

type ByteSize int

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a)
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b)
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c)
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d)
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e)
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f)

	fmt.Printf("%d \t\t\t %b \n", KB, KB)
	fmt.Printf("%d \t\t\t %b\n", MB, MB)
	fmt.Printf("%d \t\t\t %b\n", GB, GB)
	fmt.Printf("%d \t\t\t %b\n", TB, TB)
	fmt.Printf("%d \t\t\t %b\n", PB, PB)
	fmt.Printf("%d \t\t\t %b\n", EB, EB)

}
