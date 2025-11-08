package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	x, err1 := strconv.Atoi(os.Args[1])
	y, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil || x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			isTop := i == 0
			isBottom := i == y-1
			isLeft := j == 0
			isRight := j == x-1

			if (isTop || isBottom) && (isLeft || isRight) {
				if isTop && isLeft && (!isBottom || !isRight) {
					fmt.Print("A")
				} else if isTop && isBottom && isLeft && isRight {
					fmt.Print("A")
				} else if isBottom && isRight {
					fmt.Print("C")
				} else if isTop && isRight {
					fmt.Print("C")
				} else {
					fmt.Print("A")
				}
			} else if isTop || isBottom || isLeft || isRight {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
