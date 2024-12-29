package main

import (
	"fmt"
	"os"
)

//Whenever we call a method where we are updating the value we pass in a pointer

// Create custom types
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// Format the bill (receiver functions)
func (b *bill) format() string {
	fs := "Bill breakdown :\n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v.....%v\n", k+":", v)
		total += v
	}
	// add to tip
	fs += fmt.Sprintf("%-25v.....%.2f\n", "tip:", b.tip)
	//total
	fs += fmt.Sprintf("%-25v....%0.2f", "total:", total+b.tip)

	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an Item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill
func (b *bill) saveBill() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Saved Bill :", b.name)
}
