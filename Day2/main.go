package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main(){
	filePath := "input.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read lines from the file
	scanner := bufio.NewScanner(file)
	var sum int
	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text()

	
}