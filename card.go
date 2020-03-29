package main

import (
	"errors"
	"strconv"
)

var maxBalance = 90
var minBalance = 0

func negativeValueCheck(n int) error {
	if n < 0 {
		return errors.New("Cannot use negative value")
	}
	return nil
}

type card struct {
	balance int
}

func (c *card) topup(n int) error {
	err := negativeValueCheck(n)
	if err != nil {
		return err
	}

	nb := c.balance + n
	if nb > maxBalance {
		return errors.New("Cannot exceed maximum balance of " + strconv.Itoa(maxBalance))
	}
	c.balance = nb
	return nil
}

func (c *card) deduct(n int) error {
	err := negativeValueCheck(n)
	if err != nil {
		return err
	}

	nb := c.balance - n
	if nb < minBalance {
		return errors.New("Cannot reduce below minimum balance of " + strconv.Itoa(minBalance))
	}
	c.balance = nb
	return nil
}
