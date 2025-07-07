package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func main() {

	var crimes int = 0
	var dangerousSuspectName string
	var suspects = make(map[string]Man)
	suspects["Ivan"] = Man{"Ivan", "Ivanov", 23, "male", 0}
	suspects["Vasiliy"] = Man{"Vasiliy", "Petrov", 36, "male", 0}
	suspects["Nataliya"] = Man{"Nataliya", "Sidorova", 18, "female", 5}
	suspects["Mark"] = Man{"Mark", "Kharitonov", 57, "male", 0}
	suspects["Kirill"] = Man{"Kirill", "Golubev", 18, "male", 0}
	suspects["Anatoliy"] = Man{"Anatoliy", "Medinskiy", 55, "male", 12}
	suspects["Varvara"] = Man{"Varvara", "Tatischeva", 75, "female", 0}
	suspects["Petr"] = Man{"Petr", "Petrov", 33, "male", 0}
	suspects["Vasilisa"] = Man{"Vasilisa", "Bochkareva", 25, "female", 0}
	suspects["Boris"] = Man{"Boris", "Lastochkin", 19, "male", 0}

	for name, suspect := range suspects {
		if suspect.Crimes > 0 {
			if suspect.Crimes > crimes {
				crimes = suspect.Crimes
				dangerousSuspectName = name
			}
		}

	}

	if crimes > 0 {
		fmt.Printf(
			"Самый опасный подозреваемый!\nИмя:%s\nФамилия:%s\nВозраст:%d\nПол:%s\nКоличество преступлений:%d\n",
			suspects[dangerousSuspectName].Name,
			suspects[dangerousSuspectName].LastName,
			suspects[dangerousSuspectName].Age,
			suspects[dangerousSuspectName].Gender,
			suspects[dangerousSuspectName].Crimes,
		)
	} else {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
	}

}
