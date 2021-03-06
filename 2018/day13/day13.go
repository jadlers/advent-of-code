package main

import (
	"fmt"
	util "github.com/jadlers/advent-of-code/util"
	"sort"
	"strings"
)

func main() {
	lines := util.ReadLines()
	p1, p2 := Day12(lines)

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

type Position struct {
	x, y int
}

type Cart struct {
	direction        string
	nextIntersection string
}

func Day12(lines []string) (p1, p2 string) {
	// Read input
	carts := map[Position]Cart{}
	tracks := make([][]string, len(lines)) // tracks[y][x]
	for y, row := range lines {
		tracks[y] = []string{}
		for x, char := range strings.Split(row, "") {
			switch char {
			case ">":
				fallthrough
			case "<":
				tracks[y] = append(tracks[y], "-")
				carts[Position{x, y}] = Cart{char, "left"}
			case "^":
				fallthrough
			case "v":
				tracks[y] = append(tracks[y], "|")
				carts[Position{x, y}] = Cart{char, "left"}
			default:
				tracks[y] = append(tracks[y], char)
			}
		}
	}

	for len(carts) > 1 {
		_, carts, p1 = oneTick(tracks, carts)
	}

	for pos, _ := range carts {
		p2 = fmt.Sprintf("%v,%v\n", pos.x, pos.y)
	}

	return
}

func oneTick(tracks [][]string, carts map[Position]Cart) (bool, map[Position]Cart, string) {
	firstCrash := ""
	crashFound := false
	// Sort carts on tracks, only use the alive ones
	cartPositions := []Position{}
	for position := range carts {
		cartPositions = append(cartPositions, position)
	}
	sort.SliceStable(cartPositions, func(i, j int) bool {
		if cartPositions[i].x != cartPositions[j].x {
			return cartPositions[i].x < cartPositions[j].x
		}
		return cartPositions[i].y < cartPositions[j].y
	})

	// fmt.Println(cartPositions)

	// Move every cart one step
	for _, pos := range cartPositions {
		// fmt.Printf("Next cart at pos: %v: %v\n", pos, carts[pos])
		curCart := carts[pos]
		delete(carts, pos)
		switch curCart.direction {
		case ">":
			nextTrack := tracks[pos.y][pos.x+1]
			if hasCart(pos.x+1, pos.y, carts) {
				firstCrash = fmt.Sprintf("%v,%v", pos.x+1, pos.y)
				delete(carts, Position{pos.x + 1, pos.y})
				crashFound = true
				break
			} else if nextTrack == "\\" {
				curCart.direction = "v"
			} else if nextTrack == "/" {
				curCart.direction = "^"
			} else if nextTrack == "+" {
				curCart.direction, curCart.nextIntersection = getIntersectionDirection(curCart.direction, curCart.nextIntersection)
			} else if nextTrack != "-" {
				fmt.Println("ERROR >, nextTrack:", nextTrack)
			}
			if !hasCart(pos.x+1, pos.y, carts) {
				carts[Position{pos.x + 1, pos.y}] = curCart
			}
		case "v":
			nextTrack := tracks[pos.y+1][pos.x]
			if hasCart(pos.x, pos.y+1, carts) {
				firstCrash = fmt.Sprintf("%v,%v", pos.x, pos.y+1)
				delete(carts, Position{pos.x, pos.y + 1})
				crashFound = true
				break
			} else if nextTrack == "\\" {
				curCart.direction = ">"
			} else if nextTrack == "/" {
				curCart.direction = "<"
			} else if nextTrack == "+" {
				curCart.direction, curCart.nextIntersection = getIntersectionDirection(curCart.direction, curCart.nextIntersection)
			} else if nextTrack != "|" {
				fmt.Println("ERROR v, next:", nextTrack)
			}
			if !hasCart(pos.x, pos.y+1, carts) {
				carts[Position{pos.x, pos.y + 1}] = curCart
			}
		case "^":
			nextTrack := tracks[pos.y-1][pos.x]
			if hasCart(pos.x, pos.y-1, carts) {
				firstCrash = fmt.Sprintf("%v,%v", pos.x, pos.y-1)
				delete(carts, Position{pos.x, pos.y - 1})
				crashFound = true
				break
			} else if nextTrack == "\\" {
				curCart.direction = "<"
			} else if nextTrack == "/" {
				curCart.direction = ">"
			} else if nextTrack == "+" {
				curCart.direction, curCart.nextIntersection = getIntersectionDirection(curCart.direction, curCart.nextIntersection)
			} else if nextTrack != "|" {
				fmt.Println("ERROR >, nextTrack:", nextTrack)
			}
			if !hasCart(pos.x, pos.y-1, carts) {
				carts[Position{pos.x, pos.y - 1}] = curCart
			}
		case "<":
			nextTrack := tracks[pos.y][pos.x-1]
			if hasCart(pos.x-1, pos.y, carts) {
				firstCrash = fmt.Sprintf("%v,%v", pos.x-1, pos.y)
				delete(carts, Position{pos.x - 1, pos.y})
				crashFound = true
				break
			} else if nextTrack == "\\" {
				curCart.direction = "^"
			} else if nextTrack == "/" {
				curCart.direction = "v"
			} else if nextTrack == "+" {
				curCart.direction, curCart.nextIntersection = getIntersectionDirection(curCart.direction, curCart.nextIntersection)
			} else if nextTrack != "-" {
				fmt.Println("ERROR >, nextTrack:", nextTrack)
			}
			if !hasCart(pos.x-1, pos.y, carts) {
				carts[Position{pos.x - 1, pos.y}] = curCart
			}
		}
	}
	// fmt.Println(i)

	return crashFound, carts, firstCrash
}

func hasCart(x, y int, carts map[Position]Cart) bool {
	if _, cart := carts[Position{x, y}]; cart {
		return true
	}
	return false
}

func getIntersectionDirection(dir, next string) (string, string) {
	if next == "straight" {
		return dir, "right"
	} else if next == "left" {
		switch dir {
		case "<":
			return "v", "straight"
		case ">":
			return "^", "straight"
		case "^":
			return "<", "straight"
		case "v":
			return ">", "straight"
		}
	} else if next == "right" {
		switch dir {
		case "<":
			return "^", "left"
		case ">":
			return "v", "left"
		case "^":
			return ">", "left"
		case "v":
			return "<", "left"
		}
	}
	fmt.Printf("Error, no valid intersection (dir: %v, next: %v)\n", dir, next)
	return "", ""
}

func PrintMap(tracks [][]string, carts map[Position]Cart) {
	for y, row := range tracks {
		for x, char := range row {
			if _, cart := carts[Position{x, y}]; cart {
				fmt.Print(carts[Position{x, y}].direction)
			} else {
				fmt.Print(char)
			}
		}
		fmt.Println()
	}
}
