package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	filePath := "C:/Users/John/Downloads/values.txt"

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
		var numberString string
		line = strings.Replace(line, "one", "o1e", -1)
		line = strings.Replace(line, "two", "t2o", -1)
		line = strings.Replace(line, "three", "t3e", -1)
		line = strings.Replace(line, "four", "4r", -1)
		line = strings.Replace(line, "five", "5e", -1)
		line = strings.Replace(line, "six", "6x", -1)
		line = strings.Replace(line, "seven", "7n", -1)
		line = strings.Replace(line, "eight", "e8t", -1)
		line = strings.Replace(line, "nine", "n9e", -1)

		for _, char := range line {
			if unicode.IsDigit(char) {
				// Found a digit, print it
				numberString = numberString + string(char)
				//fmt.Printf("Found digit in first order: %c\n", char)
				// For the first digit found, you can break the loop
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			char := rune(line[i]) // Convert byte to rune
			if unicode.IsDigit(char) {
				// Found a digit, print it
				//fmt.Printf("Found digit in reverse order: %c\n", char)
				numberString = numberString + string(char)

				// For the last digit found, you can break the loop
				break
			}
		}
		fmt.Printf("Adding Number: %s\n", numberString)
		number, err := strconv.Atoi(numberString)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			continue // Skip to the next iteration in case of an error
		}
		sum += number

	}
	fmt.Printf("Final Running Sum: %d\n", sum)

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}
