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

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canIDrink2(age int) bool {
	// switch {
	// case age < 18:
	// 	return false
	// case age == 18:
	// 	return true
	// }
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge == 18:
		return true

	}
	return false
}

type person struct {
	name    string
	age     int
	favFood []string
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

	fmt.Println(canIDrink2(16))

	a := 2
	b := &a
	a = 10
	fmt.Println(&a, b)
	fmt.Println(*b)

	names1 := [5]string{"1", "2", "3", "4", "5"}
	names2 := []string{"nico", "lynn", "dal", "marl", "flynn"}
	names2 = append(names2, "lalalal")
	fmt.Println(names1)
	fmt.Println(names2)

	Lee := map[string]string{"name": "L", "age": "12"}
	fmt.Println(Lee)

	for key, value := range Lee {
		fmt.Println(key, value)
	}

	favFood := []string{"kimchi", "ramen"}
	l := person{"Lee", 25, favFood}
	i := person{name: "Lee", age: 18, favFood: favFood}
	fmt.Println(l)
	fmt.Println(i)
}
