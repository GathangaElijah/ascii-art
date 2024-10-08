package asciimodifier

import (
	"strings"
)


func AsciiGraphics(inputString string, asciimap map[rune][]string) string {
	// var data string
	var output strings.Builder
	var temp [][]string
	splitString := strings.Split(inputString, "\\n")
	for _, words := range splitString {
		if len(words) == 0 {
			output.WriteString("\n")
			continue
		}
		for _, char := range words {
			if asciiRep, ok := asciimap[char]; ok {
				temp = append(temp, asciiRep)
			}
		}
		
		for i := 0; i < 8; i++ {
			for j := 0; j < len(words); j++ {
				output.WriteString(temp[j][i])
			}
			if i < 8 {
				output.WriteString("\n")
			}
		}
		temp = nil
	}

	return output.String()
}
