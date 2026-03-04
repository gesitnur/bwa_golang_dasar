package main

import (
	"fmt"
	"pertama/calculation"
	"pertama/calculation_two"
)

func main() {
	fmt.Println("Hello, World!")

	result := calculation.Add(10, 20)
	fmt.Println(result)

	result_two := calculation_two.Multiply(10, 20)
	fmt.Println(result_two)

	age := 20
	fmt.Printf("Umur saya adalah %d tahun\n", age)

	if age >= 18 {
		fmt.Println("Anda sudah dewasa")
	} else {
		fmt.Println("Anda masih anak-anak")
	}

	// switch
	number := 2
	switch number {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	default:
		fmt.Println("Angka tidak dikenal")
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("Perulangan ke-%d\n", i)
	}

	title := "Golang Programming"
	for index, char := range title {
		fmt.Printf("Karakter ke-%d: %c\n", index, char)
	}

	title2 := "Golang the best language"
	for index, char := range title2 {
		if index%2 == 0 {
			fmt.Printf("Karakter ke-%d: %c\n", index, char)

		}
	}

	title3 := "Golang the best language"
	for index, char := range title3 {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			fmt.Printf("Karakter ke-%d: %c adalah huruf vokal\n", index, char)
		default:
			fmt.Printf("Karakter ke-%d: %c adalah huruf konsonan\n", index, char)
		}
	}

	var languages [5]string
	languages[0] = "Go"
	languages[1] = "Python"
	languages[2] = "JavaScript"
	languages[3] = "Java"
	languages[4] = "C++"

	fmt.Println("Bahasa pemrograman yang saya kuasai:")

	languages2 := [5]string{"Go", "Python", "JavaScript", "Java", "C++"}
	fmt.Println(languages2)
}
