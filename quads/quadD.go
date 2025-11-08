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
				// quadD: TL=A, TR=C, BL=C, BR=A
				if isTop && isLeft && !isRight {
					fmt.Print("A") // TL only
				} else if isTop && isBottom && isLeft && isRight {
					fmt.Print("A") // Single cell (1x1)
				} else if isTop && isRight && !isBottom {
					fmt.Print("C") // TR only
				} else if isBottom && isRight && !isTop {
					fmt.Print("A") // BR only
				} else if isBottom && isLeft && !isTop {
					fmt.Print("C") // BL only
				} else if isTop && isRight {
					fmt.Print("C") // TR (single row, right corner)
				} else if isBottom && isLeft {
					fmt.Print("C") // BL (single column, bottom corner)
				} else {
					fmt.Print("A") // TL (single column/row, left or top corner)
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
