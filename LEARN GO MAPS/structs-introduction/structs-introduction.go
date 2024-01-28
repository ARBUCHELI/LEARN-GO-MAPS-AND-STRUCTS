package main

import "fmt"

type Pet struct {
	name    string
	petType string
	age     int
}

func main() {
	nuggets := Pet{"Nuggets", "dog", 4}
	mittens := Pet{"Mittens", "cat", 7}
	robin := Pet{"Robin", "bird", 2}
	fmt.Println(nuggets)
	fmt.Println(mittens)
	fmt.Println(robin)
}