package main_test

import (
	"github.com/gogunit/gunit/hammy"
	"github.com/mlmon/surveyor"
	"testing"
)

func Test_which_fails_with_non_existent_binary(t *testing.T) {
	assert := hammy.New(t)
	hasBinary := main.Which("plants-and-bees-knees")
	assert.Is(hammy.False(hasBinary))
}

func Test_which_succeeds_with_sh_binary(t *testing.T) {
	assert := hammy.New(t)
	hasBinary := main.Which("sh")
	assert.Is(hammy.True(hasBinary))
}

func stubFalseyWhich() func() {
	oldWhich := main.Which
	main.Which = func(binary string) bool { return false }
	return func() { main.Which = oldWhich }
}

func stubTruthyWhich() func() {
	oldWhich := main.Which
	main.Which = func(binary string) bool {
		return true
	}
	return func() { main.Which = oldWhich }
}
