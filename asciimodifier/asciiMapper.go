package asciimodifier

import (
	"bufio"
	"fmt"
	"os"
)

// This function maps runes between 32 and 126 to their
// ascii graphical representations.


func AsciiMapper(filename string) map[rune][]string{
	// Opening the file
 file, err := os.Open(filename)
	if err != nil{
		fmt.Println("Error opening file.\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Defining the variables to be used
	
	var asciiMap = make(map[rune][]string)
	var runeValue rune = 32
	// Creating a scanner object
	scanner := bufio.NewScanner(file)
	var lines []string

	count := 0
	for scanner.Scan(){
		lineData := scanner.Text()
		
		// removing the new lines
		if len(lineData) == 0{
			continue
		} 
		lines = append(lines, lineData)
		count++
		if count == 8{
			asciiMap[runeValue] = lines
			runeValue++
			count = 0
			lines = nil
		}
		
		if runeValue > 126{
			break
		}
	}
	if err := scanner.Err(); err != nil{
		fmt.Print("Error scanning the file\n", err)
		os.Exit(1)
	}
	return asciiMap
}
