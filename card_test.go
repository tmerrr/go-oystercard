package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardInit(t *testing.T) {
	c := card{balance: 10, isInJourney: true}
	assert.Equal(t, c.balance, 10, "expected balance to equal 10")
	assert.True(t, c.isInJourney, "expected isInJourney to be true")
	assert.IsType(t, c.journeys, []journey{}, "expected journeys to be slice of type journey")
}

func TestNewCard(t *testing.T) {
	c := newCard(10)
	assert.IsType(t, card{}, c, "expected type of card")
	assert.Equal(t, c.balance, 10, "expected balance to equal 10")
	assert.False(t, c.isInJourney, "expected isInJourney to be false")
	assert.Len(t, c.journeys, 0, "expected journeys to have length of 0")
}

func TestTopup(t *testing.T) {
	c := card{}
	err := c.topup(5)
	assert.Nil(t, err, "expected err to be nil")
	assert.Equal(t, c.balance, 5, "expected balance to equal 5")
}

func TestTopupMaxLimit(t *testing.T) {
	c := newCard(90)
	err := c.topup(1)
	assert.EqualError(t, err, "Cannot exceed maximum balance of 90", "expected max balance error")
	assert.Equal(t, c.balance, 90, "expected balance to be 90")
}

func TestTopupNegativeValue(t *testing.T) {
	c := card{}
	err := c.topup(-1)
	assert.EqualError(t, err, "Cannot use negative value", "expected negative value error")
	assert.Equal(t, c.balance, 0, "expected balance to equal 0")
}

func TestDeduct(t *testing.T) {
	c := newCard(10)
	err := c.deduct(2)
	assert.Nil(t, err, "expected err to be nil")
	assert.Equal(t, c.balance, 8, "expected balance to equal 8")
}

func TestDeductMinLimit(t *testing.T) {
	c := card{}
	err := c.deduct(1)
	assert.EqualErrorf(t, err, "Cannot reduce below minimum balance of 0", "expected min balance error")
	assert.Equal(t, c.balance, 0, "expected balance to be 0")
}

func TestDeductNegativeValue(t *testing.T) {
	c := newCard(5)
	err := c.deduct(-1)
	assert.EqualError(t, err, "Cannot use negative value", "expected negative value error")
	assert.Equal(t, c.balance, 5, "expected balance to equal 5")
}

func TestTapIn(t *testing.T) {
	c := newCard(1)
	s := station{"Old St", 1}
	err := c.tapIn(s)
	assert.Nil(t, err, "expected err to be nil")
	assert.True(t, c.isInJourney, "expected isInJourney to be true")
	assert.Len(t, c.journeys, 1, "expected journeys to have length of 1")
	assert.Equal(t, c.journeys[0].start, s, "expected journey start to equal correct station")
}

func TestTapInMinBalance(t *testing.T) {
	c := card{}
	s := station{"Old St", 1}
	err := c.tapIn(s)
	assert.EqualError(t, err, "Insufficient funds. Must have minimum balance of 1", "expected insufficient funds error")
	assert.False(t, c.isInJourney, "expected isInJourney to be false")
	assert.Len(t, c.journeys, 0, "expected journeys to have length of 0")
}

func TestTapOut(t *testing.T) {
	c := newCard(1)
	// set up current journey
	c.tapIn(station{"Bank", 1})

	s := station{"London Bridge", 1}
	c.tapOut(s)
	assert.False(t, c.isInJourney, "expected isInJourney to be false")
	assert.Equal(t, c.balance, 0, "expected balance to be 0")
	assert.Len(t, c.journeys, 1, "expected journeys to have length of 1")
	assert.Equal(t, c.journeys[0].end, s, "expected journey end to equal correct station")
}
