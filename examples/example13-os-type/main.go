package main

import (
	"fmt"
	"runtime"
)

func GetOSType() string {
	switch os := runtime.GOOS; os {
	case "darwin": 
		return "osx";
	case "linux":
		return "linux"
	case "windows":
		return "windows"
	}
	return "undefined"
}

func main() {
	fmt.Println("Os type is:", GetOSType())
}