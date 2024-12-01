// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// type RGB struct {
// 	red   int
// 	green int
// 	blue  int
// 	name  int
// }

// func initRGB() *RGB {
// 	return &RGB{
// 		red:   0,
// 		green: 0,
// 		blue:  0,
// 		name:  0,
// 	}
// }

// func (r *RGB) validate(color string, value int) {
// 	if color == "red" {
// 		if value > r.red {
// 			r.red = value
// 		}
// 	} else if color == "blue" {
// 		if value > r.blue {
// 			r.blue = value
// 		}
// 	} else if color == "green" {
// 		if value > r.green {
// 			r.green = value
// 		}
// 	}
// }

// func main() {
// 	// data structure to take a set, then to see the max number
// 	// take numbers up to
// 	allRGB := []*RGB{}

// 	bytes, _ := os.ReadFile("adventday2.txt")
// 	allGames := strings.Split(string(bytes), "\n")

// 	for _, game := range allGames {
// 		rgb := initRGB()

// 		gameAndValues := strings.Split(game, ":")

// 		rgb.name, _ = strconv.Atoi(gameAndValues[0][5:])
// 		for _, roll := range strings.Split(gameAndValues[1], ";") {
// 			values := strings.Split(roll, ",")
// 			for _, value := range values {
// 				if strings.Contains(value, "red") {
// 					newRed, _ := strconv.Atoi(strings.TrimSpace(strings.Split(value, " ")[1]))
// 					rgb.validate("red", newRed)
// 				} else if strings.Contains(value, "green") {
// 					newGreen, _ := strconv.Atoi(strings.TrimSpace(strings.Split(value, " ")[1]))
// 					rgb.validate("green", newGreen)
// 				} else if strings.Contains(value, "blue") {
// 					newBlue, _ := strconv.Atoi(strings.TrimSpace(strings.Split(value, " ")[1]))
// 					rgb.validate("blue", newBlue)
// 				}
// 			}

// 		}
// 		allRGB = append(allRGB, rgb)
// 	}

// 	gameCounter := 0
// 	for _, rgb := range allRGB {
// 		if rgb.red <= 12 &&
// 			rgb.green <= 13 &&
// 			rgb.blue <= 14 {
// 			gameCounter += rgb.name
// 		}
// 	}

// 	fmt.Println(gameCounter)

// }

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RGB struct {
	red   int
	green int
	blue  int
	name  int
}

func initRGB() *RGB {
	return &RGB{
		red:   0,
		green: 0,
		blue:  0,
		name:  0,
	}
}

func (r *RGB) validate(color string, value int) {
	if color == "red" {
		if value > r.red {
			r.red = value
		}
	} else if color == "blue" {
		if value > r.blue {
			r.blue = value
		}
	} else if color == "green" {
		if value > r.green {
			r.green = value
		}
	}
}

func main() {
	// data structure to take a set, then to see the max number
	// take numbers up to
	allRGB := []*RGB{}

	bytes, _ := os.ReadFile("adventday2.txt")
	allGames := strings.Split(string(bytes), "\n")

	for _, game := range allGames {
		rgb := initRGB()

		gameAndValues := strings.Split(game, ":")

		rgb.name, _ = strconv.Atoi(gameAndValues[0][5:])
		for _, roll := range strings.Split(gameAndValues[1], ";") {
			values := strings.Split(roll, ",")
			for _, value := range values {
				if strings.Contains(value, "red") {
					newRed, _ := strconv.Atoi(strings.TrimSpace(strings.Split(value, " ")[1]))
					rgb.validate("red", newRed)
				} else if strings.Contains(value, "green") {
					newGreen, _ := strconv.Atoi(strings.TrimSpace(strings.Split(value, " ")[1]))
					rgb.validate("green", newGreen)
				} else if strings.Contains(value, "blue") {
					newBlue, _ := strconv.Atoi(strings.TrimSpace(strings.Split(value, " ")[1]))
					rgb.validate("blue", newBlue)
				}
			}

		}
		allRGB = append(allRGB, rgb)
	}

	gameCounter := 0
	for _, rgb := range allRGB {
		gameCounter += rgb.red * rgb.blue * rgb.green
	}

	fmt.Println(gameCounter)

}
