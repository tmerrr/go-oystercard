package main

import (
	"fmt"
	"os"
)

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}

func main() {
	s1 := station{"St Pauls", 1}
	s2 := station{"Finsbury Park", 2}

	c := newCard(10)

	fmt.Println("Tapping in at", s1.name)
	err := c.tapIn(s1)
	handleError(err)

	fmt.Println("Tapping out at", s2.name)
	c.tapOut(s2)

	fmt.Println(c.journeyLog.history())
	fmt.Println("Balance:", c.balance)

	fmt.Println("Tapping out at", s1.name)
	c.tapOut(s1)

	fmt.Println(c.journeyLog.history())
	fmt.Println("Balance:", c.balance)

	fmt.Println("Tapping in at", s1.name)
	err = c.tapIn(s1)
	handleError(err)

	fmt.Println("Tapping out at", s2.name)
	c.tapOut(s2)

	fmt.Println(c.journeyLog.history())
	fmt.Println("Balance:", c.balance)

	fmt.Println("Tapping in at", s1.name)
	err = c.tapIn(s1)
	handleError(err)
}
