package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	main()
}

func TestBlinkTimes(t *testing.T) {
	list := map[int]int{125: 1, 17: 1}
	times := blinkTimes(list, 6)
	assert.Equal(t, 22, times)

	list = map[int]int{125: 1, 17: 1}
	times = blinkTimes(list, 25)
	assert.Equal(t, 55312, times)
}
