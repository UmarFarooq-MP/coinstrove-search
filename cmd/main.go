package main

import "fmt"

//byte ,int , float , string
// |    | | | | | | |
// 0/1 0/1
//int (number without decimal , 3 ,4 , 12321) 32bit(4 byte)
//float (decimal numbers. 2.3 , 2.32343243) (8 bytes) -> 64 bit
//boolean (true/false (0/1)). 1 byte (8 bits)
//byte (8bit) 1 byte.  = 255 (unsigned)

func main() {
	//var age int

	//age = 143

	//var ageInNaima float64 //(8 bytes)
	//ageInNaima = 12.31
	//
	//ageInNaima = 12.4445555
	//fmt.Println(ageInNaima)
	//
	//var naimaBool bool
	//naimaBool = true
	//naimaBool = false
	//

	var i int8 = 120
	i += 10
	fmt.Println(i)

}
