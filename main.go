package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Create a new bills name: ")
	//name, _ := reader.ReadString('\n')
	//name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bills name: ", reader)

	b := newBill(name)
	fmt.Println("Created new bills :", b.name)
	return b
}

func propmtOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose option {a- add Item , s- save bills ,t - add tip }: ", reader)
	switch option {
	case "a":
		name, _ := getInput("Add  name: ", reader)
		price, _ := getInput("Add  price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Invalid price")
			propmtOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Added bills :", name, " with price :", price)
		propmtOptions(b)
	case "s":
		fmt.Println("Saving bill : ", b.format())
		b.saveBill()
	case "t":
		tip, _ := getInput("Add  tip: ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Invalid tip")
			propmtOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Updated bills with tip  :", "$", tip)
		propmtOptions(b)
	default:
		fmt.Println("Invalid Option")
		propmtOptions(b)
	}
}

func main() {
	myBill := createBill()
	propmtOptions(myBill)
}
