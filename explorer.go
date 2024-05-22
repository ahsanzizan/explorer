package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func HandleCommands(reader *bufio.Reader, currentPath string) {
	for {
		fmt.Printf("Current Directory: %s\n", currentPath)
		listDirectory(currentPath)

		fmt.Print("\n> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		command := strings.Split(input, " ")

		switch command[0] {
		case "cd":
			changeDirectory(&currentPath, command)
		case "open":
			openFile(currentPath, command)
		case "delete":
			deleteFile(currentPath, command)
		case "mkdir":
			makeDir(currentPath, command)
		case "help":
			showHelp()
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command. Type 'help' to list available commands.")
		}
		fmt.Println()
	}
}

func listDirectory(currentPath string) {
	files, err := os.ReadDir(currentPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("[DIR]  %s\n", file.Name())
		} else {
			fmt.Printf("[FILE] %s\n", file.Name())
		}
	}
}

func changeDirectory(currentPath *string, command []string) {
	if len(command) < 2 {
		fmt.Println("Usage: cd <directory>")
		return
	}
	newPath := filepath.Join(*currentPath, command[1])
	if _, err := os.Stat(newPath); os.IsNotExist(err) {
		fmt.Println("Directory does not exist.")
	} else {
		*currentPath = newPath
	}
}

func openFile(currentPath string, command []string) {
	if len(command) < 2 {
		fmt.Println("Usage: open <file>")
		return
	}
	filePath := filepath.Join(currentPath, command[1])
	if fileContent, err := os.ReadFile(filePath); err != nil {
		fmt.Println("Error reading file:", err)
	} else {
		fmt.Println(string(fileContent))
	}
}

func deleteFile(currentPath string, command []string) {
	if len(command) < 2 {
		fmt.Println("Usage: delete <file>")
		return
	}
	filePath := filepath.Join(currentPath, command[1])
	if err := os.Remove(filePath); err != nil {
		fmt.Println("Error deleting file:", err)
	} else {
		fmt.Println("File deleted.")
	}
}

func makeDir(currentPath string, command []string) {
	if len(command) < 2 {
		fmt.Println("Usage: mkdir <directory>")
		return
	}
	dirPath := filepath.Join(currentPath, command[1])
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		fmt.Println("Error creating directory:", err)
	} else {
		fmt.Println("Directory created.")
	}
}

func showHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  cd <directory>   - Change to the specified directory")
	fmt.Println("  open <file>      - Open and display the content of the specified file")
	fmt.Println("  delete <file>    - Delete the specified file")
	fmt.Println("  mkdir <directory>- Create a new directory")
	fmt.Println("  help             - Show this help message")
	fmt.Println("  exit             - Exit the file explorer")
}
