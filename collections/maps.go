package main

import "fmt"

func MapsDemo() {

	fmt.Println("Maps Demo:")

	//basic map declaration

	var userAges map[string]int
	fmt.Println("Empty map (nil):", userAges)

	//init map with make()

	userAges = make(map[string]int)
	userAges["Yusuf"] = 24
	userAges["Kaya"] = 30
	fmt.Println("User Ages:", userAges)

	//map literal
	scores := map[string]int{
		"Yusuf": 6,
		"Ahmet": 7,
		"Mehmet": 5,
	}

	fmt.Println("Scores:", scores)

	//accessing value

	fmt.Println("Yusuf's score:", scores["Yusuf"])

	//check key is exist

	value, exists := scores["Zeynep"]

	if exists {
		fmt.Println("Zeynep's score is", value)
	} else {
		fmt.Println("Zeynep not found")
	}

	//update and delete

	scores["Mehmet"] = 10
	delete(scores, "Ahmet")
	fmt.Println("Updated scores", scores)

	//loop through map

	for name, score := range scores{
		fmt.Printf("%s -> %d\n", name, score)
	}

	//map length

	fmt.Println("Total Entries:", len(scores))

	//map is reference type (shared memory)
	newScores := scores
	newScores["Yusuf"] = 100
	fmt.Println("Original map:", scores)
	fmt.Println("New map reference:", newScores)

	//map of slices
	friends := map[string][]string{
		"Alice": {"Bob", "Charlie"},
		"Bob":   {"Alice"},
	}
	fmt.Println("Map of slices:", friends)

}