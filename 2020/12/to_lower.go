package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	string :=scanner.Text()
    	splitString := strings.SplitN(string, ",", 2)
    	id := splitString[0]
    	stringsToCapitalize := strings.Split(splitString[1], " ");
    	for i, s := range stringsToCapitalize {
    		stringsToCapitalize[i] = strings.Title(s)
    	}

    	fmt.Println(id, "!~~!", strings.ReplaceAll(strings.Join(stringsToCapitalize[:], " "), "\"", "\\\""))
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}