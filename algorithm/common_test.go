package algorithm

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

func TestGetSingleton(t *testing.T) {
	s1 := GetSingleton()
	s2 := GetSingleton()

	if s1 != s2 {
		t.Fatal("not equal")
	}

	assert.Equal(t, s1, s2)
}

func TestDecorator(t *testing.T) {
	coolFunc(myFunc)
}

func TestGetIota(t *testing.T) {
	GetIota()
}

func TestFac(t *testing.T) {
	fmt.Println(Sum(4))
}
func TestPriFib(t *testing.T) {
	PriFib(5)
}