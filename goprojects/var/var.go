package main

import "fmt"

func var1() {
	var a int = 10
	var msg string = "Hello Variable"

	a = 20
	msg = "Good Morning"
	fmt.Println(msg, a)
}

func var2() {
	var minWage int = 10
	var workingHour int = 20

	var income int = minWage * workingHour

	fmt.Println(minWage, workingHour, income)
}

func main() {
	var1()
	var2()
}
