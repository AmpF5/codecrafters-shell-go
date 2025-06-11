package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var _ = fmt.Fprint

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		query, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}

		command := strings.Fields(query[:len(query)-1])

		commandLen := len(command)

		if commandLen == 0 {
			fmt.Println("No command entered")
			continue
		}

		if command[0] == "exit" {
			if commandLen == 1 {
				fmt.Println("Missing parameter")
				continue
			}

			if command[1] == "0" {
				os.Exit(0)
			} else {
				fmt.Println("Invalid parameter, use 0 to exit")
			}
		}
	}
}
