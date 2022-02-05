package warm

import (
	"reflect"
	"testing"
)

	func TestGetUniqelementSlice(t *testing.T) {

		// test 1
		input1 := []string{"a", "b", "c", "d"}
		input2 := []string{"a", "b", "c", "d"}

		got := GetUniqeElementSlice(input1, input2)

		want := []string{"a", "b", "c", "d"}

		if reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)

		}

		// test 2
		input1 = []string{"a", "b", "c", "d"}
		input2 = []string{"b", "c", "d"}
		input3 := []string{"a", "d", "z"}

		want = []string{"a", "b", "c", "d", "z"}

		got = GetUniqeElementSlice(input1, input2, input3)

		if reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)

		}

	}
