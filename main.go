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

	// Chosing which data file to display depending on what the user wants
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

	//Validating the files
	fileData := asciimodifier.FileReader(bannerName)
	currentHash := asciimodifier.CalculateHash(fileData)
	// fmt.Println("This is the current hash =>", currentHash)
	var originalHash = make(map[string]string)
	originalHash["standard"] = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	originalHash["shadow"] = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	originalHash["thinkertoy"] = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
	// fmt.Println("This is the length of the map =>", len(originalHash))
	if bannerHash, ok := originalHash[fileName]; ok{
		if bannerHash != currentHash{
			fmt.Println("Error: file modified")
			return
		}
	} else{
		fmt.Println("No such file")
		return
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
