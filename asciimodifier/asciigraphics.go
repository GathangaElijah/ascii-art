package asciimodifier

// import "fmt"

func AsciiGraphics(inputString string, asciimap map[rune][]string) string {
	// var data string
	var temp [][]string
	for _, char := range inputString {
		temp = append(temp, asciimap[char])
	}
	word := ""
	for i := 0; i < 8; i ++ {
		for j := 0; j < len(inputString); j++ {
			word += temp[j][i]
		}
		if i != 7 {
			word+= "\n"
		}


	}
	// for _, char := range inputString {
	// 	if graphicRep, ok := asciimap[char]; ok {
	// 		fmt.Println(string(graphicRep))
	// 		// data +=  string(graphicRep)+"\n"
	// 	}
	// }


	return word
}
