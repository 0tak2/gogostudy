package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func Minus(a int, b int) int {
	return a - b
}

func Operate(op int, a int, b int) int {
	if op == 1 {
		return Add(a, b)
	} else if op == 2 {
		return Minus(a, b)
	} else {
		fmt.Println("Invalid Op Code")
		return -1
	}
}

func main() {
	var op, a, b, result int

	fmt.Println("Usage: OPERATION_TYPE A B")
	fmt.Println("OPERATION_TYPE: 1(Add), 2(Minus)")

	_, err := fmt.Scanf("%d %d %d", &op, &a, &b)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		result = Operate(op, a, b)
	}
	fmt.Println(result)
}
