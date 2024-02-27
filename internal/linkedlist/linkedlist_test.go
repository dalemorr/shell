package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	val := 5
	list := new(LinkedList[int])
	list.Add(val)

	out, err := list.Get(0)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, val, out)
}
