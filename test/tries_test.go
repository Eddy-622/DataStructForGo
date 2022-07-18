package test

import (
	"DataStructures/data_struct"
	"testing"
)

func TestTries(t *testing.T) {
	tries := data_struct.NewTries()
	tries.Insert("cat")
	tries.Insert("can")

	t.Log('b' == rune(98))
}
