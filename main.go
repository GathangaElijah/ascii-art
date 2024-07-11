package main

import (
	"fmt"
	"os"

	"ascii-art/asciimodifier"
)

func main() {
	// Reading the input from the commandline
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: go run . [string] [standard/thinkertoy/shadow]")
		os.Exit(0)
	}

	// Chosing whic data file to display depending on what the user wants
	bannerName := "bannerfiles/standard.txt"
	fileName := os.Args[2]
	if fileName != "standard" && fileName != "shadow" && fileName != "thinkertoy" {
		fmt.Println("Usage:pass either [standard/thinkertoy/shadow]")
		return
	}

	switch fileName {
	case "thinkertoy":
		bannerName = "bannerfiles/thinkertoy.txt"
	case "shadow":
		bannerName = "bannerfiles/shadow.txt"
	default:
		bannerName = "bannerfiles/standard.txt"
	}

	//Handling the input string
	var inputString = os.Args[1]
	if inputString == "\\n"{
		fmt.Print("\n")
		return
	} else if inputString == ""{
		return
	}

	asciiMap := asciimodifier.AsciiMapper(bannerName)
	// fmt.Println(asciiMap)
	
	asciiGraphics := asciimodifier.AsciiGraphics(inputString, asciiMap)
	fmt.Print(asciiGraphics)
}
