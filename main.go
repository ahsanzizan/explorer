package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadUser(reader *bufio.Reader, currentPath string) {
	for {
		fmt.Printf("Current Directory: %s\n", currentPath)
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

		fmt.Print("\n> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		command := strings.Split(input, " ")

		switch command[0] {
		case "cd":
			if len(command) < 2 {
				fmt.Println("Usage: cd <directory>")
			} else {
				newPath := filepath.Join(currentPath, command[1])
				if _, err := os.Stat(newPath); os.IsNotExist(err) {
					fmt.Println("Directory does not exist.")
				} else {
					currentPath = newPath
				}
			}
		case "open":
			if len(command) < 2 {
				fmt.Println("Usage: open <file>")
			} else {
				filePath := filepath.Join(currentPath, command[1])
				if fileContent, err := os.ReadFile(filePath); err != nil {
					fmt.Println("Error reading file:", err)
				} else {
					fmt.Println(string(fileContent))
				}
			}
		case "delete":
			if len(command) < 2 {
				fmt.Println("Usage: delete <file>")
			} else {
				filePath := filepath.Join(currentPath, command[1])
				if err := os.Remove(filePath); err != nil {
					fmt.Println("Error deleting file:", err)
				} else {
					fmt.Println("File deleted.")
				}
			}
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command. Available commands: cd, open, delete, exit")
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	ReadUser(reader, currentPath)
}
