package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

var cache = make(map[string]int)

func printHeader() {
	fmt.Println("Welcome To Coding Dojo - Cache Entry (11/27/2020)\n")
	fmt.Println("Type 'm' for manual\n")
}

func getInput(message string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(message)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

    return input
}

func isManualRequest(input string) bool {
	return input == "m"
}

func printManual() {
	fmt.Println("------------------")
	fmt.Println("Cache Entry Manual")
	fmt.Println("------------------")
	fmt.Println("")
	fmt.Println("m \t  print manual")
	fmt.Println("c <entry> clear entry")
	fmt.Println("l <entry> list entry count")
	fmt.Println("C \t  clear all entries")
	fmt.Println("q \t  quit")
	fmt.Println("")
}

func shouldQuit(input string) bool {
	return input == "q"
}

func isClearEntry(input string) bool {
	data := strings.Split(input, " ")
	return len(data) == 2 && data[0] == "c"
}

func clearEntry(input string) {
	cache[input] = 0
}

func isListEntry(input string) bool {
	data := strings.Split(input, " ")
	return len(data) == 2 && data[0] == "l"	
}

func listEntry(input string) {
	fmt.Println(cache[input])
}

func isClearAllEntries(input string) bool {
	return input == "C"
}

func clearAllEntries() {
	cache = make(map[string]int)
}

func addToCache(input string) {
    count, ok := cache[input]
    if ok {
    	cache[input] = count + 1
    } else {
    	cache[input] = 1
    }
}

func runLoop() {
	for {
		input := getInput(">> ")

		if shouldQuit(input) {
			break
		} else if isManualRequest(input) {
			printManual()
		} else if isClearEntry(input) {
			clearEntry(strings.Split(input, " ")[1])
		} else if isListEntry(input) {
			listEntry(strings.Split(input, " ")[1])
		} else if isClearAllEntries(input) {
			clearAllEntries()
		} else {
			addToCache(input)
		}
	}
}

func main() {
	printHeader()
	runLoop()
}