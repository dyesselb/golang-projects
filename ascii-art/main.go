package main

import (
	"fmt"
	"os"
	funcs "student/functions"
)

func main() {
	if !funcs.CheckFile("standard.txt") {
		fmt.Println("ERROR: standard.txt is not correct!")
		return
	}
	if len(os.Args) == 2 {
		if os.Args[1] == "" {
			return
		}
		funcs.CreateArt(os.Args[1])
	} else {
		fmt.Println("ERROR: You must enter only one argument!")
	}
}
