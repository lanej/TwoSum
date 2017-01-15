package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	expected := 3
	actual := solve(load("pa611.txt"), -10000, 10000)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	expected := 5
	actual := solve(load("pa612.txt"), -10000, 10000)

	assert.Equal(t, expected, actual)
}
