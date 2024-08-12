package asciimodifier

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func FileReader(filename string) []byte {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error \n", err)
		os.Exit(1)
	}
return fileData
}

func CalculateHash(filedata []byte) string {
fileHash := sha256.Sum256(filedata)
return fmt.Sprintf("%x", fileHash)
}
