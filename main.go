package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	flag "github.com/ogier/pflag"
)

var (
	userInput string
)

func main() {
	flag.Parse()
	if !isValidFlags() {
		printUsage()
		os.Exit(1)
	}
	fetchGitUsers()
}

func init() {
	flag.StringVarP(&userInput, "user", "u", "", "Search users")
}

func isValidFlags() bool {
	if flag.NFlag() == 0 {
		return false
	}
	return true
}

func printUsage() {
	fmt.Printf("Usage: %s [options] \n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func fetchGitUsers() {
	users := strings.Split(userInput, ",")
	fmt.Printf("Searching for user(s): %s \n", users)

	for _, user := range users {
		if gitUser, err := getUserFromGit(user); err == nil {
			pp, _ := json.MarshalIndent(gitUser, "", "    ")
			fmt.Printf("%s \n", string(pp))
			fmt.Printf("\n")
		} else {
			log.Fatal(err)
		}
	}
}
