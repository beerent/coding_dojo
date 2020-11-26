package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func getNumberOfRounds() (rounds int) {
	roundsQuestion := "Best of how many rounds?"
    rounds = getIntegerInput(roundsQuestion)
    for {
    	if rounds % 2 == 1 {
    		break
    	}
    	fmt.Println("must be an odd number of rounds.")

    	rounds = getIntegerInput(roundsQuestion)
    }

	return
}

func getIntegerInput(message string) int {
	reader := bufio.NewReader(os.Stdin)

    var i int

    for {
    	fmt.Print(message, " ")
    	input, _ := reader.ReadString('\n')
    	input = strings.TrimSuffix(input, "\n")
    	intVar, err := strconv.Atoi(input)

    	if (err == nil) {
    		i = intVar
    		break
    	} else {
    		fmt.Println("invalid number '" + input + "'")
    	}
    }

	return i
}

func main() {
    fmt.Println("Welcome To Rock Paper Scissors Shoot (11/26/2020)")
    fmt.Println("")
    rounds := getNumberOfRounds()

    fmt.Println(rounds)
}