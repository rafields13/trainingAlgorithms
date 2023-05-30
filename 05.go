package main

import "fmt"

func main() {
	emails := []string{"alex", "steven", "alex", "steven", "alex", "steven", "alex", "steven"}
	registeredEmails := make(map[string]int)

	for _, email := range emails {
		count, ok := registeredEmails[email]

		if !ok {
			registeredEmails[email] = 1
			fmt.Println("OK")
		} else {
			newEmail := fmt.Sprintf("%s%d", email, count)
			registeredEmails[email]++
			registeredEmails[newEmail] = 1
			fmt.Println(newEmail)
		}
	}
}
