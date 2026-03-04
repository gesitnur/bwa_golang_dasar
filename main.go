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
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("Perulangan ke-%d\n", i)
	// }

	// title := "Golang Programming"
	// for index, char := range title {
	// 	fmt.Printf("Karakter ke-%d: %c\n", index, char)
	// }

	// title2 := "Golang the best language"
	// for index, char := range title2 {
	// 	if index%2 == 0 {
	// 		fmt.Printf("Karakter ke-%d: %c\n", index, char)

	// 	}
	// }

	// title3 := "Golang the best language"
	// for index, char := range title3 {
	// 	switch char {
	// 	case 'a', 'e', 'i', 'o', 'u':
	// 		fmt.Printf("Karakter ke-%d: %c adalah huruf vokal\n", index, char)
	// 	default:
	// 		fmt.Printf("Karakter ke-%d: %c adalah huruf konsonan\n", index, char)
	// 	}
	// }

	// var languages [5]string
	// languages[0] = "Go"
	// languages[1] = "Python"
	// languages[2] = "JavaScript"
	// languages[3] = "Java"
	// languages[4] = "C++"

	// fmt.Println("Bahasa pemrograman yang saya kuasai:")

	// languages2 := [5]string{"Go", "Python", "JavaScript", "Java", "C++"}
	// fmt.Println(languages2)

	// // slice
	// var gamingConsole []string
	// gamingConsole = append(gamingConsole, "PlayStation 5")
	// // gamingConsole = append(gamingConsole, "Xbox Series X")
	// // gamingConsole = append(gamingConsole, "Nintendo Switch")

	// fmt.Println(gamingConsole)

	// // with value
	// gamingConsole2 := []string{"PlayStation 5", "Xbox Series X", "Nintendo Switch"}
	// fmt.Println(gamingConsole2)

	// map
	// var myMap map[string]int
	// // myMap = make(map[string]int)
	// myMap = map[string]int{}

	// myMap["satu"] = 1
	// myMap["dua"] = 2
	// myMap["tiga"] = 3

	// fmt.Println(myMap)

	// myMap2 := map[string]int{
	// 	"satu": 1,
	// 	"dua":  2,
	// 	"tiga": 3,
	// }

	// fmt.Println(myMap2)

	// for key, value := range myMap2 {
	// 	fmt.Printf("Key: %s, Value: %d\n", key, value)
	// }

	// delete(myMap2, "dua")
	// fmt.Println(myMap2)

	// value, isExist := myMap2["dua"]
	// if isExist {
	// 	fmt.Printf("Nilai dari key 'tiga' adalah %d\n", value)
	// } else {
	// 	fmt.Println("Key 'tiga' tidak ditemukan")
	// }

	// // Slice of map
	// // slice dari sebuah map yang keynya string dan value nya string
	// students := []map[string]string{
	// 	{
	// 		"name": "John Doe",
	// 		"age":  "20",
	// 	},
	// 	{
	// 		"name": "Jane Smith",
	// 		"age":  "22",
	// 	},
	// }

	// fmt.Println(students)

	// for _, student := range students {
	// 	fmt.Printf("Nama: %s, Umur: %s\n", student["name"], student["age"])
	// }

	// scores := [8]int{90, 85, 78, 92, 88, 95, 80, 91}
	// var total int
	// for _, score := range scores {
	// 	total += score
	// }
	// average := float64(total) / float64(len(scores))
	// fmt.Printf("Rata-rata nilai: %.2f\n", average)

	scores := [8]int{90, 85, 78, 92, 88, 95, 80, 91}
	var goodScores []int
	for index, score := range scores {
		if score >= 90 {
			goodScores = append(goodScores, score)
			fmt.Printf("Nilai ke-%d: %d adalah nilai yang baik\n", index, score)
		}
	}
	fmt.Printf("Nilai yang baik: %v\n", goodScores)

}
