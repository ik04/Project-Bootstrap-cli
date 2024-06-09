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
	if _, err := os.Stat(folder); !os.IsNotExist(err) {
		color.Red("Folder %s already exists. Please choose a different name.", folder)
		os.Exit(1)
	}

	runCommand("mkdir", folder)
	if err := os.Chdir(folder); err != nil {
		panic(err)
	}
}

func main() {
	color.Blue("Choose setup:")
	options := map[int64]string{
		1: "Nextjs with Laravel",
		2: "Remixjs with Laravel",
		3: "Nextjs with Nodejs",
	}

	for k, v := range options {
		optionStr := fmt.Sprintf("%v) %v", k, v)
		color.Cyan(optionStr)
	}

	reader := bufio.NewReader(os.Stdin)

	optionStr, _ := getInput("Select your option: ", reader)
	option, err := strconv.ParseInt(optionStr, 10, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Selected:", options[option])

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting current directory:", err)
		return
	}

	var projectName string
	var projectPath string
	for {
		projectName, _ = getInput("Enter Project Name: ", reader)
		projectPath = fmt.Sprintf("%s/%s", currentDir, projectName)
		if _, err := os.Stat(projectPath); os.IsNotExist(err) {
			break
		}
		color.Red("Folder %s already exists. Please choose a different name.", projectPath)
	}

	switch option {
	case 1:
		makeAndCd(projectName)
		makeAndCd("client")
		runCommand("npx", "create-next-app@latest", "web")
		color.Blue("Nextjs Setup Finished!")
		if err := os.Chdir(projectPath); err != nil {
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
		if err := os.Chdir(projectPath); err != nil {
			fmt.Fprintln(os.Stderr, "Error changing directory:", err)
			return
		}
		makeAndCd("server")
		runCommand("laravel", "new", "rest")
		color.Blue("Laravel Setup Finished!")
		color.Green("Project Setup Finished!")

	case 3:
		makeAndCd(projectName)
		makeAndCd("client")
		runCommand("npx", "create-next-app@latest", "web")
		color.Blue("Nextjs Setup Finished!")
		if err := os.Chdir(projectPath); err != nil {
			fmt.Fprintln(os.Stderr, "Error changing directory:", err)
			return
		}
		makeAndCd("server")
		runCommand("pnpm", "init", "-y")
		color.Blue("Node Setup Finished!")
		color.Green("Project Setup Finished!")

	default:
		fmt.Println("Invalid option")
	}
}
