package main

import (
	"fmt"

	"github.com/fernok/dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	dictionary["hello"] = "hello"

	// Search
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	// Add
	baseword := "hi"
	errAdd := dictionary.Add(baseword, "Greeting")
	if errAdd != nil {
		fmt.Println(errAdd)
	}
	definitionAdd, _ := dictionary.Search("hi")
	fmt.Println(definitionAdd)

	// Update
	errUpdate := dictionary.Update(baseword, "hello")
	if errUpdate != nil {
		fmt.Println(errUpdate)
	}
	updatedWord, _ := dictionary.Search(baseword)
	fmt.Println(updatedWord)

	// Delete
	dictionary.Delete(baseword)
	_, errDelete := dictionary.Search(baseword)
	fmt.Println(errDelete)
}
