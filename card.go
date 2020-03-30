package main

import (
	"errors"
	"strconv"
)

// declare constants
var (
	maxBalance  = 90
	minBalance  = 0
	minFare     = 1
	penaltyFare = 6
)

func negativeValueCheck(n int) error {
	if n < 0 {
		return errors.New("Cannot use negative value")
	}
	return nil
}

type card struct {
	balance     int
	isInJourney bool
	journeys    []journey
}

func newCard(balance int) card {
	return card{balance, false, []journey{}}
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

func (c *card) tapIn(s station) error {
	if c.balance < minFare {
		return errors.New("Insufficient funds. Must have minimum balance of 1")
	}
	c.journeys = append(c.journeys, journey{start: &s})
	c.isInJourney = true
	return nil
}

func (c *card) tapOut(s station) {
	i := len(c.journeys) - 1
	j := c.journeys[i]
	j.end = &s
	c.journeys[i] = j
	c.balance = c.balance - minFare
	c.isInJourney = false
}
