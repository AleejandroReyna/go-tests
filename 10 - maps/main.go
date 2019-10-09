package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["es"] = "Spanish"
	languages["en"] = "English"

	otherLanguages := map[string]string{"fr": "France", "ch": "Chinese"}
	numberKeys := map[int]string{1: "One", 2: "Two"}

	fmt.Println(languages)
	fmt.Println(otherLanguages)

	delete(otherLanguages, "fr")

	fmt.Println(otherLanguages)

	if language, ok := languages["es"]; ok {
		fmt.Println(language)
	} else {
		fmt.Println("Language not exist.")
	}

	fmt.Println(numberKeys)
}
