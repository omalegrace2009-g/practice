package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) ([]string, error) {
	ban, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	banner := strings.Split(string(ban), "\n")
	return banner, err
}

func GetLine(banner []string, c rune) []string {
	if c < 32 || c > 126 {
		fmt.Println("Error: Input not in banner file!!")
		return nil
	}
	start := (int(c) - 32) * 9
	return banner[start+1 : start+9]
}

func PrintArt(banner []string, print []string) {
	for _, play := range print {
		if play == "" {
			fmt.Println()
			continue
		}
		for i := 0; i <= 7; i++ {
			for _, let := range play {
				get := GetLine(banner, let)
				if get == nil {
					return
				} else {
					fmt.Print(get[i])
				}
			}
			fmt.Println()
		}
	}
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=<color> <substring to be colored> something")
		return
	}
	// run := []string{
	// 	"red",
	// 	"blue",
	// 	"black",
	// 	"pink",
	// 	"purple",
	// 	"orange",
	// 	"green",
	// }
	input := os.Args[1]
	if strings.HasPrefix(input, "--color=") {
 		g := input[8:]
		input2 := os.Args[2]
	input3 := os.Args[3]
	}
	line := strings.Split(input, "\\n")
	if line[0] == "" {
		line = line[1:]
	}
	ban, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println("Error reading banner file", err)
		return
	}
	if len(os.Args) == 3 {
		input1 := os.Args[2] + ".txt"
		ben, err := LoadBanner(input1)
		if err != nil {
			fmt.Println("Error reading banner file", err)
			return
		}
		PrintArt(ben, line)
	} else {
		PrintArt(ban, line)
	}
}
