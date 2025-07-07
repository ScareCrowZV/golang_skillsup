package main

import "fmt"

func rangeOverString(s string) {
	for i, item := range s {
		fmt.Println(i, fmt.Sprintf("%#U", item))
	}
}

func main() {
	str := "©"

	fmt.Println(len(str))
	fmt.Println(str)

	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}

	fmt.Println("\nВыводим через Range")
	rangeOverString(str)

	runes := []rune(str)
	fmt.Println(fmt.Sprintf("%#U", runes[0]))

}
