package main

import "fmt"
import "os"

func loadingFile(fileName string) ([]byte, error) {
	text, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error loading file: ", err)
		return nil, err
	}
	return text, nil
}

func main() {
	content, err := loadingFile("data.txt")
	if err != nil {
		fmt.Println("Error loading file ", err)
		return
	}
	fmt.Println("The file content is", string(content))
}
