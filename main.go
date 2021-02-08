package main

import (
	"fmt"
	"strings"
)

var flag bool = true

func multyply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpper2(name string) (lenght int, uppercase string) {
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func lenAndUpper3(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	fmt.Println(numbers)
	total := 0

	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return total
	//range
}

func main() {
	name := "nico"
	name = "Lynn"
	fmt.Println(name)
	fmt.Println(multyply(2, 2))

	totalLength, upperName := lenAndUpper("Lee")
	fmt.Println(totalLength, upperName)

	totalLength, upperName = lenAndUpper2("Lee")
	fmt.Println(totalLength, upperName)

	totalLength, upperName = lenAndUpper3("Lee")
	fmt.Println(totalLength, upperName)

	repeatMe("nico", "lynn", "dal", "marl", "flynn")

	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
