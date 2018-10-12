package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 3, Sum(1, 2))
}

func TestSub(t *testing.T) {
	assert.Equal(t, 3, Sub(5, 2))
}
