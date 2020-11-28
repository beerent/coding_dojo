package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
    "math/rand"
    "time"
)

func printHeader() {
	fmt.Println("Welcome To Coding Dojo - Rock Paper Scissors Shoot (11/26/2020)")
}

func printRoundsHeader(rounds int) {
	if (rounds == 1) {
		fmt.Println("\nBest of 1 round. Good luck!")
	} else {
		fmt.Println("\nBest of", rounds, "rounds. Good luck!")
	}
}

func printWinner(winner string) {
	if winner == "user" {
		fmt.Println("\nCongratulations! You are the winner.")
	} else {
		fmt.Println("\nYou lost! :(")
	}
}

func getStringInput(message string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(message, " ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	return input
}

func getIntegerInput(message string) int {
	var i int

    for {
    	input := getStringInput(message)
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

func getFullAttackName(shortname string) (fullname string) {
	switch shortname {
		case "r":
			fullname = "rock"
		case "p":
			fullname =  "paper"
		case "s":
			fullname = "scissors"
		}

	return
}

func getNumberOfRounds() int {
	roundsQuestion := "\nBest of how many rounds?"

	var rounds int
    for {
        rounds = getIntegerInput(roundsQuestion)

    	if rounds % 2 == 1 {
    		break
    	}
    	fmt.Println("must be an odd number of rounds.")
    }

	return rounds
}

func getUserDecision() string {
	decisionQuestion := "\nRock (r), Paper (p), or Scissors (s)?"
	
	var userDecision string
	for {
		userDecision = strings.ToLower(getStringInput(decisionQuestion))
		if ! isValidUserDecision(userDecision) {
			fmt.Println("invalid option '" + userDecision + "'")
			userDecision = getStringInput(decisionQuestion)
		}

		break
	}

	return userDecision
}

func isValidUserDecision(userDecision string) bool {
	return userDecision == "r" || userDecision == "p" || userDecision == "s"
}

func getComputerDecision() string {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(3)

	if randomNumber == 0 {
		return "r"
	} else if randomNumber == 1 {
		return "p"
	}

	return "s"
}

func getRoundWinner(userDecision, computerDecision string) string {
	switch userDecision {
	case "r":
		switch computerDecision {
		case "s":
			return "user"
		case "p":
			return "computer"
		}
		break

	case "p":
		switch computerDecision {
		case "r":
			return "user"
		case "s":
			return "computer"
		}
		break

	case "s":
		switch computerDecision {
		case "p":
			return "user"
		case "r":
			return "computer"
		}
		break
	}

	return "tie"
}

func printRoundResults(userDecision, computerDecision, roundWinner string) {
	fmt.Println("User throws\t", getFullAttackName(userDecision))
	fmt.Println("Computer throws\t", getFullAttackName(computerDecision))
	fmt.Println(roundWinner + "!")
}

func getWinner(userWins, computerWins, rounds int) (winner string, gameover bool) {
	winsNeeded := (rounds + 1) / 2

	if userWins >= winsNeeded {
		return "user", true
	}

	if computerWins >= winsNeeded {
		return "computer", true
	}

	return "", false
}

func winnerIsUser(winner string) bool {
	return winner == "user"
}

func winnerIsComputer(winner string) bool {
	return winner == "computer"
}

func runGame(rounds int) {
	printRoundsHeader(rounds);

	userWins := 0
	computerWins := 0

	var winner string
	
	for {
		userDecision := getUserDecision()
		computerDecision := getComputerDecision()
		roundWinner := getRoundWinner(userDecision, computerDecision)
		printRoundResults(userDecision, computerDecision, roundWinner)

		if winnerIsUser(roundWinner) {
			userWins++
		} else if winnerIsComputer(roundWinner) {
			computerWins++
		}

		winnerName, gameover := getWinner(userWins, computerWins, rounds)
		if gameover {
			winner = winnerName
			break
		}
	}

	printWinner(winner)
}

func main() {
    printHeader()
    runGame(getNumberOfRounds())
}