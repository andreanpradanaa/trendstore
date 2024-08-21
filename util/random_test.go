package util

import (
	"fmt"
	"testing"
)

func TestRandomDescription(t *testing.T) {
	n := 10

	desc := RandomDescription(n)

	fmt.Println(desc)
}
