package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// getInput function to read user input from the console
func getInput(prompt string, r *bufio.Reader) (string, error) {
	color.HiCyan(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

// runCommand function to run shell commands
func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin 
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

// makeAndCd function to create a directory and change into it
func makeAndCd(folder string) {
	// Check if folder exists
	if _, err := os.Stat(folder); !os.IsNotExist(err) {
		color.Red("Folder %s already exists. Please choose a different name.", folder)
		os.Exit(1)
	}

	// Create folder and change into it
	runCommand("mkdir", folder)
	if err := os.Chdir(folder); err != nil {
		panic(err)
	}
}

// main function
func main() {
	color.Blue("Choose setup:")
	options := map[int64]string{
		1: "Nextjs with Laravel",
		2: "Remixjs with Laravel",
	}

	// Display options
	for k, v := range options {
		optionStr := fmt.Sprintf("%v) %v", k, v)
		color.Cyan(optionStr)
	}

	// Create a new reader
	reader := bufio.NewReader(os.Stdin)

	// Read user option
	optionStr, _ := getInput("Select your option: ", reader)
	option, err := strconv.ParseInt(optionStr, 10, 64)
	if err != nil {
		panic(err)
	}

	// Print selected option
	fmt.Println("Selected:", options[option])

	// Read project name
	var projectName string
	for {
		projectName, _ = getInput("Enter Project Name: ", reader)
		if _, err := os.Stat(projectName); os.IsNotExist(err) {
			break
		}
		color.Red("Folder %s already exists. Please choose a different name.", projectName)
	}

	switch option {
	case 1:
		// Change to home directory
		if err := os.Chdir(os.Getenv("HOME")); err != nil {
			fmt.Fprintln(os.Stderr, "Error changing directory:", err)
			return
		}

		// Setup Nextjs and Laravel
		makeAndCd(projectName)
		makeAndCd("client")
		runCommand("npx", "create-next-app@latest", "web")
		color.Blue("Nextjs Setup Finished!")
		rootPath := fmt.Sprintf("%v/%v", os.Getenv("HOME"), projectName)
		if err := os.Chdir(rootPath); err != nil {
			fmt.Fprintln(os.Stderr, "Error changing directory:", err)
			return
		}
		makeAndCd("server")
		runCommand("laravel", "new", "rest")
		color.Blue("Laravel Setup Finished!")
		color.Green("Project Setup Finished!")

	case 2:
		makeAndCd(projectName)
		makeAndCd("client")
		runCommand("npx", "create-remix@latest", "web")
		color.Blue("Remixjs Setup Finished!")
		makeAndCd("../server")
		runCommand("laravel", "new", "rest")
		color.Blue("Laravel Setup Finished!")
		color.Green("Project Setup Finished!")

	default:
		fmt.Println("Invalid option")
	}
}
