package main

import (
	"errors"
	"strconv"
)

var maxBalance = 90

type card struct {
	balance int
}

func (c *card) topup(n int) error {
	nb := c.balance + n
	if nb > maxBalance {
		return errors.New("Cannot exceed maximum balance of " + strconv.Itoa(maxBalance))
	}
	c.balance = nb
	return nil
}
