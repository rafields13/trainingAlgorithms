package main

import "fmt"

type Participant struct {
	Name string
	Role string
}

func maximumNumberTeams(participants []Participant) int {
	numbersMathematicians := 0

	for _, participant := range participants {
		if participant.Role == "Mathematician" {
			numbersMathematicians++
		}
	}
	numbersProgrammers := 0

	for _, participant := range participants {
		if participant.Role == "Programmer" {
			numbersProgrammers++
		}
	}

	if len(participants) < 4 {
		return 0
	} else if len(participants) >= 4 && numbersMathematicians == 0 || numbersProgrammers == 0 {
		return 0
	} else if len(participants) == 4 {
		return 1
	} else if len(participants) > 4 && numbersMathematicians < len(participants)/4 {
		return numbersMathematicians
	} else if len(participants) > 4 && numbersProgrammers < len(participants)/4 {
		return numbersProgrammers
	}

	return len(participants) / 4
}

func main() {
	participants := []Participant{
		{
			Name: "Vicky",
			Role: "Programmer",
		},
		{
			Name: "John",
			Role: "Programmer",
		},
		{
			Name: "Luis",
			Role: "Programmer",
		},
		{
			Name: "Lucas",
			Role: "Mathematician",
		},
		{
			Name: "July",
			Role: "Mathematician",
		},
		{
			Name: "Mary",
			Role: "Mathematician",
		},
		{
			Name: "Sam",
			Role: "Mathematician",
		},
		{
			Name: "Joshua",
			Role: "Mathematician",
		},
		{
			Name: "Fred",
			Role: "Mathematician",
		},
		{
			Name: "Andrew",
			Role: "Mathematician",
		},
		{
			Name: "Leo",
			Role: "Mathematician",
		},
		{
			Name: "Gabriel",
			Role: "Mathematician",
		},
		{
			Name: "Anne",
			Role: "Mathematician",
		},
		{
			Name: "Steven",
			Role: "Mathematician",
		},
		{
			Name: "Mathew",
			Role: "Mathematician",
		},
		{
			Name: "Edward",
			Role: "Mathematician",
		},
		{
			Name: "Iza",
			Role: "Mathematician",
		},
		{
			Name: "Raphael",
			Role: "Mathematician",
		},
	}

	showing := maximumNumberTeams(participants)
	fmt.Print(showing)
}
