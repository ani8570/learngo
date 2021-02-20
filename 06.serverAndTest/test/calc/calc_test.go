package calc_test

import (
	"testing"

	"github.com/ani8570/learngo/06.serverAndTest/test/calc"
)

func TestSum(t *testing.T) {
	s := calc.Sum(1, 2, 3)

	if s != 6 {
		t.Error("Wrong result")
	}
}
