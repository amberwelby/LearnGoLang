package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

type menu []menuItem

// Method
func (m menu) print() {
	for _, item := range m {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, cost := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, cost)
		}
	}
}

func (m *menu) add() error {
	fmt.Println("Please enter the name of the new item")
	name, _ := in.ReadString('\n')
	name = strings.TrimSpace(name)
	for _, item := range data {
		if item.name == name {
			return errors.New("menu item already exists")
		}
	}
	*m = append(*m, menuItem{name: name, prices: make(map[string]float64)})
	return nil // Returned with no error
}

var in = bufio.NewReader(os.Stdin)

// Functions
func AddItem() error {
	return data.add()
}

func PrintMenu() {
	data.print()
}
