package main

import "fmt"

func mapLoop() {
	v := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"give":  5,
	}

	delete(v, "give")

	for key, value := range v {
		fmt.Println(key, value)
	}
}

func main() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
	if value, ok := x["vasya"]; ok {
		fmt.Println(value, ok)
	}

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"

	fmt.Println(elements["H"])

	elements2 := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
	}

	fmt.Println(elements2["He"])

	system := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "gas",
		},
		"Be": {
			"name":  "Beryllium",
			"state": "gas",
		},
	}

	fmt.Println(system["Li"])

	mapLoop()

}
