package test

import (
	"regexp"
	"testing"
)

func TestName(t *testing.T) {
	ok, err := regexp.MatchString(`^[ABCD]\d\d$`, "AA10")
	t.Log(ok, err)
}
