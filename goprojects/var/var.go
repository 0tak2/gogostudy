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

func var3() {
	var a int = 3
	var b int // 초기값 생략. 타입별 기본값으로 초기화됨
	var c = 4 // 타입 생략. 우변의 값으로부터 추론된 타입이 지정됨 (정수는 int, 실수는 float64)
	d := 5    // 선언 대입문 := 사용. var 키워드와 타입 생략 가능

	fmt.Println(a, b, c, d)

	/*
		타입별 기본값

		정수 타입 (intN, uintN, ...)	0
		실수 타입 (floatN, complexN)	0.0
		불리언							false
		문자열							""
		그 외							nil
	*/
}

func main() {
	var1()
	var2()
	var3()
}
