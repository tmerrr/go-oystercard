package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCard(t *testing.T) {
	c := card{10}
	assert.Equal(t, c.balance, 10, "expected balance to equal 10")
}

func TestTopup(t *testing.T) {
	c := card{}
	err := c.topup(5)
	assert.Nil(t, err, "expected err to be nil")
	assert.Equal(t, c.balance, 5, "expected balance to equal 5")
}

func TestTopupMaxLimit(t *testing.T) {
	c := card{90}
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
	c := card{10}
	err := c.deduct(2)
	assert.Nil(t, err, "expected err to be nil")
	assert.Equal(t, c.balance, 8, "expected balance to equal 8")
}

func TestDeductMinLimit(t *testing.T) {
	c := card{0}
	err := c.deduct(1)
	assert.EqualErrorf(t, err, "Cannot reduce below minimum balance of 0", "expected min balance error")
	assert.Equal(t, c.balance, 0, "expected balance to be 0")
}

func TestDeductNegativeValue(t *testing.T) {
	c := card{5}
	err := c.deduct(-1)
	assert.EqualError(t, err, "Cannot use negative value", "expected negative value error")
	assert.Equal(t, c.balance, 5, "expected balance to equal 5")
}
