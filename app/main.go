package main

import (
	"bufio"
	"fmt"
	"os"
)

var _ = fmt.Fprint

func main() {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Fprintln(os.Stderr, "An error occurred:", r)
				}
			}()

			fmt.Fprint(os.Stdout, "$ ")

			query, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				panic(err)
			}

			handleCommand(query)
		}()
	}
}
