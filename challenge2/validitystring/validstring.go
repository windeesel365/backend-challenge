package validitystring

import "log"

func CheckString(s string) {
	if len(s) != 5 {
		log.Printf("Error: Please ensure to enter 5 characters")
		return
	}
	for _, char := range s {
		if char != 'L' && char != 'R' && char != '=' {
			log.Printf("Error: String contains invalid character: %c", char)
			return
		}

	}
}
