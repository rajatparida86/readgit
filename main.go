package main

import (
	"fmt"
	flag "github.com/ogier/pflag"
	"os"
	"strings"
)

var (
	userInput string
)

func main() {
	flag.Parse()
	validateInputAndFetchGitUsers()
}

func init() {
	flag.StringVarP(&userInput, "user", "u", "", "Search users")
}

func isValidFlags() {
	if flag.NFlag() == 0 {
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Printf("Usage: %s [options] \n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func validateInputAndFetchGitUsers() {
	isValidFlags()

	users := strings.Split(userInput, ",")
	fmt.Printf("Searching for user(s): %s \n", users)

	for _, user := range users {
		gitUser := getUserFromGit(user)

		fmt.Printf("Login: %s \n", gitUser.Login)
		fmt.Printf("User ID: %v \n", gitUser.ID)
		fmt.Printf("Name: %s \n", gitUser.Name)
		fmt.Printf("Url: %s \n", gitUser.URL)
		fmt.Printf("Location: %s \n", gitUser.Location)
		fmt.Printf("Email: %s \n", gitUser.Email)
		fmt.Printf("Company: %s \n", gitUser.Company)
		fmt.Printf("Bio: %s \n", gitUser.Bio)
		fmt.Printf("\n")
	}
}
