package main

import (
	"fmt"
	"math/rand"
)

var (
	Black = Color("\033[1;30m%s\033[0m") 
	Green  = Color("\033[1;32m%s\033[0m")
	Yellow = Color("\033[1;33m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func main() {
	fmt.Print("\n")
	height := 19

	for i := 0; i < height; i++ {
		for j := 0; j < height*2-1; j++ {

			if j == (height - i) {
				it := i*2 + 1
				for it != 0 {
					r := rand.Int()
					if r%10/2 > 3 {
						fmt.Print(Yellow("0"))
						it--
					} else {
						fmt.Print(Green("*"))
						it--
					}
				}

			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Print("\n")
	}
	for i := 0; i < height*2-1; i++ {
		if i == height-1 || i == height+1 {
			fmt.Printf(Black("|"))
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n\n")
}