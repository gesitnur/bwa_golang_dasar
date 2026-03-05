package main

import (
	"fmt"
)

type Student struct {
	ID        int
	FirstName string
	LastName  string
	GPA       float32
}

type Group struct {
	Name        string
	Admin       Student
	Users       []Student
	IsAvaliable bool
}

// method untuk struct
func (user Student) display() string {
	return fmt.Sprintf("ID: %d, Name display: %s", user.ID, user.FirstName)
}

func (group Group) display2() {
	fmt.Printf("Group Name: %s\n", group.Name)
	fmt.Printf("Admin Name: %s\n", group.Admin.FirstName)
	fmt.Printf("Is Available: %t\n", group.IsAvaliable)
	fmt.Println("Users:")
	for _, user := range group.Users {
		fmt.Printf("- %s %s\n", user.FirstName, user.LastName)
	}
}

func main() {
	// fmt.Println("Hello, World!")

	// result := calculation.Add(10, 20)
	// fmt.Println(result)

	// result_two := calculation_two.Multiply(10, 20)
	// fmt.Println(result_two)

	// age := 20
	// fmt.Printf("Umur saya adalah %d tahun\n", age)

	// if age >= 18 {
	// 	fmt.Println("Anda sudah dewasa")
	// } else {
	// 	fmt.Println("Anda masih anak-anak")
	// }

	// // switch
	// number := 2
	// switch number {
	// case 1:
	// 	fmt.Println("Satu")
	// case 2:
	// 	fmt.Println("Dua")
	// case 3:
	// 	fmt.Println("Tiga")
	// default:
	// 	fmt.Println("Angka tidak dikenal")
	// }

	// printMyResult("yauuuu")
	// fmt.Println(add(100, 700))
	// sum, product := calculate(10, 20)
	// fmt.Printf("Hasil penjumlahan: %d, Hasil perkalian: %d\n", sum, product)

	// fmt.Println(calculateWithPredefinedReturn(10, 20))

	// scores := []int{5, 15, 25}
	// total := sum2(scores)

	// fmt.Printf("Total skor: %d\n", total)

	// result, err := calculate3(10, 2, "+")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Printf("Hasil: %d\n", result)
	// }

	// struct
	user := Student{}
	user.ID = 1
	user.FirstName = "John"
	user.LastName = "Doe"
	fmt.Println(user.display())

	user2 := Student{
		ID:        2,
		FirstName: "Jane",
		LastName:  "Smith",
	}

	users := []Student{user, user2}

	group := Group{"Agam", user, users, true}

	displayGroup(group)
	group.display2()

	fmt.Println(group)

	fmt.Println(user.FirstName)

	fmt.Println(displayUser(user))

	// simple pointer
	numberA := 5
	numberB := &numberA

	*numberB = 10

	fmt.Println("Nilai numberA:", numberA)

	// var numberC int = 15
	// var numberD int = &numberC -> ini error
	// var numberD *int = &numberC

	// Contoh kasus penggunaan pointer
	number := 5

	change(&number, 100)
	// tanpa pointer, nilai number tidak akan berubah karena yang dikirim ke fungsi change adalah salinan dari nilai number, bukan referensi ke nilai aslinya. Dengan menggunakan pointer, kita dapat mengubah nilai number di dalam fungsi change, dan perubahan tersebut akan terlihat di luar fungsi.

	fmt.Println("Nilai number setelah fungsi change dipanggil:", number)

	// struct untuk atribut atau field dari struct lainnya

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

	// scores := [8]int{90, 85, 78, 92, 88, 95, 80, 91}
	// var goodScores []int
	// for index, score := range scores {
	// 	if score >= 90 {
	// 		goodScores = append(goodScores, score)
	// 		fmt.Printf("Nilai ke-%d: %d adalah nilai yang baik\n", index, score)
	// 	}
	// }
	// fmt.Printf("Nilai yang baik: %v\n", goodScores)

	// pointer struct sebagai parameter
	student := Student{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		GPA:       3.5,
	}

	fmt.Println("Sebelum graduate:", student.FirstName)
	graduate(&student)
	// yang dikirim adalah memori dari struct student, atau referensing ke struct student, sehingga perubahan yang dilakukan di dalam fungsi graduate akan mempengaruhi nilai dari struct student di luar fungsi tersebut. Jika kita tidak menggunakan pointer, maka perubahan yang dilakukan di dalam fungsi graduate tidak akan mempengaruhi nilai dari struct student di luar fungsi tersebut karena yang dikirim ke fungsi graduate adalah salinan dari struct student, bukan referensi ke struct student yang asli.
	fmt.Println("Setelah graduate:", student.FirstName)

}

// pointer struct sebagai parameter
func graduate(student *Student) {
	// harus menerima parameter pointer tanda bintang
	student.FirstName = "Sabar"

}

func change(number *int, newValue int) {
	*number = newValue
}

// embedded struct
func displayGroup(group Group) {
	fmt.Printf("Group Name: %s\n", group.Name)
	fmt.Printf("Admin Name: %s\n", group.Admin.FirstName)
	fmt.Printf("Is Available: %t\n", group.IsAvaliable)
	fmt.Println("Users:")
	for _, user := range group.Users {
		fmt.Printf("- %s %s\n", user.FirstName, user.LastName)
	}
}

func displayUser(user Student) string {
	result := fmt.Sprintf("ID: %d, Name: %s", user.ID, user.FirstName)
	return result
}

// function

func printMyResult(sentence string) {
	fmt.Println(sentence)
}

func add(a int, b int) int {
	return a + b
}

// return multiple value
func calculate(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

// function dengan predefined return value
func calculateWithPredefinedReturn(a int, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return
}

func sum2(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func calculate3(a int, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b != 0 {
			return a / b, nil
		} else {
			return 0, fmt.Errorf("Pembagian dengan nol tidak diperbolehkan")
		}
	default:
		return 0, fmt.Errorf("Operator tidak dikenal")
	}
}
