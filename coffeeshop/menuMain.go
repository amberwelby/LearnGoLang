package coffeeshop

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// Adding my own package
	menu "demo/coffeeshop/menu"
)

// MARK: Coffee Shop Demo App

var in = bufio.NewReader(os.Stdin)

func Operate() {
loop: // This is a label, it helps us access things like telling the switch what to break
	for {
		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) Quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.PrintMenu()
		case "2":
			menu.AddItem()
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}
}
