package main

import "fmt"

func main() {

	// var value = (((2 + 6) % 3) * 4 - 2) / 3
	// var isEqual = (value == 2)
	var left = false
	var right = true

	// fmt.Printf("nilai %d (%t) \n", value, isEqual)
	// var leftAndRight = left && right
	// fmt.Printf("left && right \t(%t) \n", leftAndRight)
	// var leftOrRight = left || right
	// fmt.Printf("left || right \t(%t) \n", leftOrRight)
	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)

	// 	Berikut penjelasan statemen operator logika pada kode di atas.
	// 	-leftAndRight bernilai false, karena hasil dari false dan true adalah false.
	// 	-leftOrRight bernilai true, karena hasil dari false atau true adalah true.
	// 	-leftReverse bernilai true, karena negasi (atau lawan dari) false adalah true.

}
