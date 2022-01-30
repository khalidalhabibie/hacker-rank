package warm

import (
	"testing"
)

func TestRepeatedString(t *testing.T) {

	var n int64
	var want int64
	// test 1
	str := "abcac"
	n = 10
	want = 4

	got := RepeatedString(str, n)
	if got != want {
		t.Errorf("error %v n %v want %v got %v", str, n, want, got)
	}

	// test 2
	str = "aba"
	n = 10
	want = 7

	got = RepeatedString(str, n)
	if got != want {
		t.Errorf("error %v n %v want %v got %v", str, n, want, got)
	}



	// test 2
	str = "a"
	n = 1000000000000
	want = n

	got = RepeatedString(str, n)
	if got != want {
		t.Errorf("error %v n %v want %v got %v", str, n, want, got)
	}
}
