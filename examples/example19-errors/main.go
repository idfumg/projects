package main

import (
	"errors"
	"fmt"
)

type FindError struct {
	Name string
	Server string
	Msg string
}

func (e FindError) Error() string {
	return fmt.Sprintf("%s", e.Msg)
}

var ErrorCrewNotFound = errors.New("Crew member not found")

var scMapping = map[string]int{
	"James": 5,
	"Kevin": 10,
	"Rahul": 9,
}

// func findSC(name string, server string) (int, error) {
// 	if value, ok := scMapping[name]; !ok {
// 		return -1, errors.New("Crew member not found")
// 	} else {
// 		return value, nil
// 	}
// }

// func findSC(name string, server string) (int, error) {
// 	if value, ok := scMapping[name]; !ok {
// 		return -1, fmt.Errorf("Crew member '%s' cound't be found on the server '%s'", name, server)
// 	} else {
// 		return value, nil
// 	}
// }

// func findSC(name string, server string) (int, error) {
// 	if value, ok := scMapping[name]; !ok {
// 		return -1, ErrorCrewNotFound
// 	} else {
// 		return value, nil
// 	}
// }

func findSC(name string, server string) (int, error) {
	if value, ok := scMapping[name]; !ok {
		return -1, FindError{name, server, "Crew member not found"}
	} else {
		return value, nil
	}
}

func main() {
	// if clearance, err := findSC("Ruko", "Server1"); err != nil {
	// 	fmt.Println("Error occured while searching for clearance level:", err)
	// } else {
	// 	fmt.Println("Clearance level found:", clearance)
	// }
	// if clearance, err := findSC("Ruko", "Server1"); err == ErrorCrewNotFound {
	// 	fmt.Println("Error occured while searching for clearance level:", err)
	// } else {
	// 	fmt.Println("Clearance level found:", clearance)
	// }
	if clearance, err := findSC("Ruko", "Server1"); err != nil {
		fmt.Println("Error occured while searching for clearance level:", err)
		if object, ok := err.(FindError); ok {
			fmt.Println("Server name:", object.Server)
			fmt.Println("Crew member name:", object.Name)
		}
	} else {
		fmt.Println("Clearance level found:", clearance)
	}
}
