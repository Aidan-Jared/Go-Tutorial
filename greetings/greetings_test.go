package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Aidan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Aidan")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Aidan") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
