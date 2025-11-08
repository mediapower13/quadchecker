package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func quadA(x, y int) []string {
	if x <= 0 || y <= 0 {
		return nil
	}
	result := make([]string, y)
	for i := 0; i < y; i++ {
		line := ""
		for j := 0; j < x; j++ {
			if (i == 0 || i == y-1) && (j == 0 || j == x-1) {
				line += "o"
			} else if i == 0 || i == y-1 {
				line += "-"
			} else if j == 0 || j == x-1 {
				line += "|"
			} else {
				line += " "
			}
		}
		result[i] = line
	}
	return result
}

func quadB(x, y int) []string {
	if x <= 0 || y <= 0 {
		return nil
	}
	result := make([]string, y)
	for i := 0; i < y; i++ {
		line := ""
		for j := 0; j < x; j++ {
			if i == 0 && j == 0 {
				line += "/"
			} else if i == 0 && j == x-1 {
				line += "\\"
			} else if i == y-1 && j == 0 {
				line += "\\"
			} else if i == y-1 && j == x-1 {
				line += "/"
			} else if i == 0 || i == y-1 {
				line += "*"
			} else if j == 0 || j == x-1 {
				line += "*"
			} else {
				line += " "
			}
		}
		result[i] = line
	}
	return result
}

func quadC(x, y int) []string {
	if x <= 0 || y <= 0 {
		return nil
	}
	result := make([]string, y)
	for i := 0; i < y; i++ {
		line := ""
		for j := 0; j < x; j++ {
			// quadC: TL=A, TR=C, BL=A, BR=C
			isTop := i == 0
			isBottom := i == y-1
			isLeft := j == 0
			isRight := j == x-1

			if (isTop || isBottom) && (isLeft || isRight) {
				// Corner - for single cell (1x1), top-left wins
				if isTop && isLeft && (!isBottom || !isRight) {
					line += "A" // TL (when not all corners overlap)
				} else if isTop && isBottom && isLeft && isRight {
					line += "A" // Single cell - TL priority
				} else if isBottom && isRight {
					line += "C" // BR
				} else if isTop && isRight {
					line += "C" // TR
				} else {
					line += "A" // BL
				}
			} else if isTop || isBottom || isLeft || isRight {
				line += "B"
			} else {
				line += " "
			}
		}
		result[i] = line
	}
	return result
}

func quadD(x, y int) []string {
	if x <= 0 || y <= 0 {
		return nil
	}
	result := make([]string, y)
	for i := 0; i < y; i++ {
		line := ""
		for j := 0; j < x; j++ {
			// quadD: TL=A, TR=C, BL=C, BR=A
			isTop := i == 0
			isBottom := i == y-1
			isLeft := j == 0
			isRight := j == x-1

			if (isTop || isBottom) && (isLeft || isRight) {
				// Corner
				if isTop && isLeft && (!isBottom || !isRight) {
					line += "A" // TL
				} else if isTop && isBottom && isLeft && isRight {
					line += "A" // Single cell
				} else if isBottom && isRight {
					line += "A" // BR
				} else if isTop && isRight {
					line += "C" // TR
				} else {
					line += "C" // BL
				}
			} else if isTop || isBottom || isLeft || isRight {
				line += "B"
			} else {
				line += " "
			}
		}
		result[i] = line
	}
	return result
}

func quadE(x, y int) []string {
	if x <= 0 || y <= 0 {
		return nil
	}
	result := make([]string, y)
	for i := 0; i < y; i++ {
		line := ""
		for j := 0; j < x; j++ {
			// quadE: TL=A, TR=C, BL=C, BR=A (same corners as quadD but different priority)
			isTop := i == 0
			isBottom := i == y-1
			isLeft := j == 0
			isRight := j == x-1

			if (isTop || isBottom) && (isLeft || isRight) {
				// Corner - quadE prioritizes left over right for bottom row
				if isTop && isLeft && (!isBottom || !isRight) {
					line += "A" // TL
				} else if isTop && isBottom && isLeft && isRight {
					line += "A" // Single cell
				} else if isBottom && isLeft {
					line += "C" // BL (left priority for bottom)
				} else if isTop && isRight {
					line += "C" // TR
				} else {
					line += "A" // BR
				}
			} else if isTop || isBottom || isLeft || isRight {
				line += "B"
			} else {
				line += " "
			}
		}
		result[i] = line
	}
	return result
}

func readInput() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func compareQuads(input []string, generated []string) bool {
	if len(input) != len(generated) {
		return false
	}
	for i := range input {
		if input[i] != generated[i] {
			return false
		}
	}
	return true
}

func main() {
	input := readInput()

	if len(input) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// Determine dimensions
	y := len(input)
	x := 0
	if y > 0 {
		x = len(input[0])
	}

	// Check if all lines have the same length
	for _, line := range input {
		if len(line) != x {
			fmt.Println("Not a quad function")
			return
		}
	}

	if x <= 0 || y <= 0 {
		fmt.Println("Not a quad function")
		return
	}

	// Check against all quad functions
	matches := []string{}

	quadFuncs := map[string]func(int, int) []string{
		"quadA": quadA,
		"quadB": quadB,
		"quadC": quadC,
		"quadD": quadD,
		"quadE": quadE,
	}

	// Check in alphabetical order
	quadNames := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}

	for _, name := range quadNames {
		generated := quadFuncs[name](x, y)
		if compareQuads(input, generated) {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", name, x, y))
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(strings.Join(matches, " || "))
	}
}
