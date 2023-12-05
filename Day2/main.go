package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//Logging
	logfile := "log.txt"

	// Open the file for writing (create if it doesn't exist, truncate if it does)
	log, err := os.Create(logfile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer log.Close()

	// Redirect standard output to the file
	oldStdout := os.Stdout
	os.Stdout = log
	defer func() { os.Stdout = oldStdout }()

	filePath := "input.txt"
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	//End Logging

	// Create a scanner to read lines from the file
	scanner := bufio.NewScanner(file)
	// Iterate over each line in the file
	var sum int
	var isvalidgame bool

	for scanner.Scan() { // This is each game

		line := scanner.Text()
		fmt.Printf("Starting a new game \n")
		fmt.Printf("New game is %s \n", line)

		//Obtain the game number first
		splitstring := strings.Split(line, ":")
		// fmt.Printf("%s is the splitstring 0 \n", splitstring[0])
		// fmt.Printf("%s is the subtstring 1 \n", splitstring[1])

		gamenum := splitstring[0][5:]
		//fmt.Printf("%s is the subtring0 \n", splitstring[0])
		fmt.Printf("%s is the gamenum \n", gamenum)
		set := strings.Split(splitstring[1], ";")
	Loop:
		for _, value := range set {
			//Evaluate each game, is it valid?
			fmt.Printf("Starting a new set \n")
			fmt.Printf("value is : %s \n", value)
			//Split the string into colors
			entry := strings.Split(value, ",")
			//Check every color
			for _, value := range entry {
				colorletter := value[len(value)-1]
				colorquant := value[:3]
				fmt.Printf("%s is the entry \n", value)
				// fmt.Printf("%s is the color letter \n", string(colorletter))
				// fmt.Printf("%s is the colorquant \n", string(colorquant))

				number, err := strconv.Atoi(strings.TrimSpace(colorquant))
				if err != nil {
					fmt.Println("Error converting string to integer: converting colorquant", err)
					continue // Skip to the next iteration in case of an error
				}

				if string(colorletter) == "d" {
					if number >= 13 {
						fmt.Printf("Breaking out of Red for game %s \n", gamenum)
						isvalidgame = false
						fmt.Printf("After failure, the sum is %d \n", sum)
						break Loop
					}
				}
				if string(colorletter) == "n" {
					if number >= 14 {
						fmt.Printf("Breaking out of Green for game %s \n", gamenum)
						isvalidgame = false
						fmt.Printf("After failure, the sum is %d \n", sum)
						break Loop
					}
				}
				if string(colorletter) == "e" {
					if number >= 15 {
						fmt.Printf("Breaking out of Blue for game %s \n", gamenum)
						isvalidgame = false
						fmt.Printf("After failure, the sum is %d \n", sum)
						break Loop
					}
				}
				fmt.Printf("Entry is a success \n")
				isvalidgame = true
			}
		}
		fmt.Printf("Set Complete!")

		if isvalidgame == true {
			toadd, err := strconv.Atoi(strings.TrimSpace(gamenum))
			if err != nil {
				fmt.Println("Error converting string to integer: mathsis hard", err)
				continue // Skip to the next iteration in case of an error
			}
			fmt.Printf("I am adding %d as it is a success case \n", toadd)
			sum += toadd
			fmt.Printf("The sum is %d \n", sum)

			isvalidgame = false
		}

	}
	fmt.Printf("final sum: %d \n", sum)

}
