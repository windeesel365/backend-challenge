package main

import (
	"bufio"
	"decodetonum/stringtonum"
	"fmt"
	"os"
	"strings"
)

func main() {

	allCases := []string{
		"LLRR=",
		"==RLL",
		"=LLRR",
		"RRL=R",
		"RLRLRL",
	}

	for _, encoded := range allCases {
		decoded := stringtonum.Stringtonum(encoded)
		fmt.Printf("Input: %s  Output: %s\n", encoded, decoded)
	}

	//เตรียมสำหรับ userInput
	var userInput string
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Please enter encoded string (only capital letters 'R', 'L', '='):")

		// read user input
		userInput, _ = reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		// validate user input จำนวนcharacters
		isValid := true
		if len(userInput) > 5 {
			isValid = false
			fmt.Println("Invalid input. Please enter only 5 characters")
		}

		// validate user input ลักษณะcharacters
		for _, char := range userInput {
			if char != 'L' && char != 'R' && char != '=' {
				isValid = false
				fmt.Println("Invalid input. Please enter only capital letters 'R', 'L', '='.")
			}

		}

		if isValid {
			// ถ้า input valid   break loop
			break
		}
	}

	//func decode
	decodedUserInput := stringtonum.Stringtonum(userInput)
	fmt.Printf("Input: %s  Output: %s\n", userInput, decodedUserInput)

}
