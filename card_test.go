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

func TestTopupLimit(t *testing.T) {
	c := card{90}
	err := c.topup(1)
	assert.NotNil(t, err, "expected err not to be nil")
	assert.Equal(t, c.balance, 90, "expected balance to be 90")
	_, isErr := err.(error)
	assert.True(t, isErr, "expected err to be of type error")
}
