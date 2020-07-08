package main

import (
	"fmt"
	"strings"
)

/*
 *	private functions start with lowercase -> cannot be exported
 *	functions starting with uppercase can be exported to another package
 */

// equivalent to multiply(a, b in)
func multiply(a int, b int) int {
	return a * b
}

// functions can return multiple values
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// ... allows the function to get unlimited arguments of that type as input
func repeatMe(words ...string) {
	fmt.Println(words)
}

// < naked return >
// length and uppercase variables are already initialized ↓ here
func lenAndUpper2(name string) (length int, uppercase string) {
	// < defer >
	// do something after the function is finished (when the function returns)
	defer fmt.Println("I'm Done")
	// not creating new length and uppercase, just updating it
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	total := 0
	// < range >
	// allows you to loop on arrays -> gives index as well
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	// before checking a condition, you can create variables
	// creating variables exclusively for the if-else statement
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canIDrink2(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	// Printing
	fmt.Println("hello, world")

	// constants and variables
	const name string = "kim" // specify type of constant
	//name = "lee" --> error! constants cannot change
	var city = "seoul"
	city = "tokyo" // variables can change
	capital := "seoul"
	/*
	 * var city = "seoul" is equivalent to city := "seoul"
	 * The shorthand version (:=), Go will guess the type for you
	 * based on the first value you assign.
	 * The shorthand version will not work outside functions.
	 */
	fmt.Println(city, capital)

	// functions part one
	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("fernok")
	fmt.Println(totalLength, upperName)
	totalLength2, _ := lenAndUpper("wildcard")
	fmt.Println(totalLength2)

	repeatMe("aaa", "bbb", "ccc", "ddd")

	// functions part two
	totLength, upperCase := lenAndUpper2("naked")
	fmt.Println(totLength, upperCase)

	// for, range, ...args
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	// if with a twist
	fmt.Println(canIDrink(16))

	// pointers
	a := 2
	b := &a
	a = 5
	fmt.Println(a, &a, b, *b)
	*b = 7
	fmt.Println(a, &a, b, *b)

	// arrays and slices
	// array: must specify length of array and type of array
	alphabets := [5]string{"a", "b", "c"}
	alphabets[3] = "d"
	alphabets[4] = "e"
	// slices: array without length
	cities := []string{"seoul", "busan", "incheon"}
	// append returns a new slice with the new element
	// takes slice and element as argument
	cities = append(cities, "inje")
	fmt.Println(cities)

	// maps
	// must specify key and value types
	fernok := map[string]string{"name": "HM", "age": "24"}
	fmt.Println(fernok)
	for key, value := range fernok {
		fmt.Println(key, value)
	}

	// structs
	favFood := []string{"kimchi", "ramen"}
	// not good coding: have to look up for reference
	nico := person{"nico", 18, favFood}
	fmt.Println(nico)
	// instead:
	nico = person{name: "nico", age: 18, favFood: favFood}
}
