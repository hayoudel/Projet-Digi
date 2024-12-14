package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 && len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go --color=<color> <string to be colored>")
		return
	}

	var colorFlag string
	var letters string

	if len(os.Args) == 3 {
		colorFlag = "--color=31"
		letters = os.Args[2]
	} else {
		colorFlag = os.Args[1]
		letters = os.Args[2]
	}

	colorCode := strings.Split(colorFlag, "=")[1]

	if len(letters) == 1 {
		fmt.Printf("\033[%sm%s\033[0m\n", colorCode, letters)
	} else {
		file, err := os.Open("art.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		art := make([]string, 256)

		for scanner.Scan() {
			line := scanner.Text()
			index := int(line[0])
			art[index] = line[2:]
		}

		for _, letter := range letters {
			index := int(letter)
			if art[index] != "" {
				lines := strings.Split(art[index], "\\n")
				for _, line := range lines {
					fmt.Printf("\033[%sm%s\033[0m\n", colorCode, line)
				}
			} else {
				fmt.Printf("\033[%sm%c\033[0m", colorCode, letter)
			}
		}
		fmt.Println()
	}
}
