package main

import (
	"electronic"
	"fmt"
	"strconv"
)

func main() {
	headerString := "====== Устройство ======"
	applePhone := electronic.NewApplePhone("iPhone 13 Pro")
	androidPhone := electronic.NewAndroidPhone("Samsung", "Galaxy S24 Ultra")
	radioPhone := electronic.NewRadioPhone("Panasonic", "KX-TG1711", 20)

	fmt.Printf("%s\n%s\n\n", headerString, printCharacteristics(applePhone))

	fmt.Printf("%s\n%s\n\n", headerString, printCharacteristics(androidPhone))

	fmt.Printf("%s\n%s\n\n", headerString, printCharacteristics(radioPhone))
}

func printCharacteristics(phone electronic.Phone) string {

	result := "Брэнд: " + phone.Brand() + ";\nМодель: " + phone.Model() + ";\nТип: " + phone.Type() + ";"

	if phone.Type() == "station" {
		result = result + "\nКоличество кнопок: " + strconv.Itoa(phone.(electronic.StationPhone).ButtonCount()) + ";"
	}

	if phone.Type() == "smartphone" {
		result = result + "\nОС: " + phone.(electronic.Smartphone).OS()
	}

	return (result)

}
