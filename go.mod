module ascii-art-color

go 1.22.2
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// var colors = map[string]string{
// 	"black":   "\033[30m",
// 	"red":     "\033[31m",
// 	"green":   "\033[32m",
// 	"yellow":  "\033[33m",
// 	"blue":    "\033[34m",
// 	"magenta": "\033[35m",
// 	"cyan":    "\033[36m",
// 	"white":   "\033[37m",
// 	"reset":   "\033[0m",
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) == 0 {
// 		return
// 	}

// 	var color, sub, text string

// 	if strings.HasPrefix(args[0], "--color=") {
// 		color = strings.TrimPrefix(args[0], "--color=")

// 		if len(args) < 2 || len(args) > 3 {
// 			return
// 		}
// 		text = args[len(args)-1]

// 		if len(args) == 3 {
// 			sub = args[1]
// 		} else {
// 			sub = text
// 		}
// 	} else {
// 		text = args[0]
// 	}

// 	code, ok := colors[color]
// if color != "" && !ok {
// 	fmt.Println("unknown color")
// 	return
// }
// if sub != "" {
// 	text = strings.ReplaceAll(
// 		text,
// 		sub,
// 		code+sub+colors["reset"],
// 	)
// }
// 	fmt.Println(strings.ReplaceAll(text, "\\n", "\n"))
// }