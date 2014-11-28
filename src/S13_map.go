package main

import "fmt"

var attended = map[string]bool{
	"Anne": false,
	"Joe":  true,
}

func main() {
	attended["Jérôme"] = false
	fmt.Println(getAttendedStatus("Jérôme"))
	fmt.Println(getAttendedStatus("John Doe"))
}

func getAttendedStatus(person string) string {
	if here, ok := attended[person]; ok { // _, ok
		if here { // puis attended[person]
			return fmt.Sprintln(person, "était à la réunion")
		} else {
			return fmt.Sprintln(person, "avait décliné l'invitation")
		}
	} else {
		return fmt.Sprintln(person, "n'était pas invité")
	}
}
