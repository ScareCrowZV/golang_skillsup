package main

import "fmt"

func main() {
	str := "℃"

	fmt.Printf("plain string: ")
	fmt.Printf("%s", str)
	fmt.Printf("\n")

	fmt.Println("представление байт в разных системах счисления")
	fmt.Println(fmt.Sprintf("десятичное: %d", str[0]))
	fmt.Println(fmt.Sprintf("двоичное: %b", str[0]))
	fmt.Println(fmt.Sprintf("шестнадцатиричное: %x", str[0]))
	fmt.Println(fmt.Sprintf("шестнадцатиричное: %x", str[1]))
	fmt.Println(fmt.Sprintf("шестнадцатиричное: %x", str[2]))
	fmt.Println(len(str))

	for i := 0; i < len(str); i++ {
		fmt.Printf("%x", str[i])
	}
}
