package main

import "fmt"

func main() {

	// var names [4]string
	// names[0] = "atrisa"
	// names[1] = "endya"
	// names[2] = "nur"
	// names[3] = "hidayah"

	// fmt.Println(names[0], names[1], names[2], names[3])

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// fmt.Println("Jumlah element \t\t", len(fruits))
	// fmt.Println("Isi semua element \t", fruits)

	// var fruits [4]string
	// // cara horizontal
	// // fruits = [4]string{"apple", "grape", "banana", "melon"}
	// // cara vertikal
	// fruits = [4]string{
	// 	"apple",
	// 	"grape",
	// 	"banana",
	// 	"melon",
	// }
	// fmt.Println("Isi semua element \t", fruits)

	// var numbers = [...]int{2, 3, 2, 4, 3}

	// fmt.Println("data array \t:", numbers)
	// fmt.Println("jumlah elemen \t:", len(numbers))

	// var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	// var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	// fmt.Println("numbers1", numbers1)
	// fmt.Println("numbers2", numbers2)

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("elemen %d : %s\n", i, fruits[i])
	// }

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for i, fruit := range fruits {
	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
	// }

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for _, fruits := range fruits {
	// 	fmt.Printf("nama buah : %s\n", fruits)
	// }

	var fruits = make([]string, 2)
	fruits[0] = "apple"
	fruits[1] = "manggo"

	fmt.Println(fruits)
}
