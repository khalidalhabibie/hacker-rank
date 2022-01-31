package warm

import "testing"

func TestJumpingOnCloud(t *testing.T) {

	dataTest1 := []int32{0, 0, 1, 0, 0, 1, 0}
	var want int32 = 4

	got := JumpingOnCloud(dataTest1)
	if got != want {
		t.Errorf("want %v and got %v", want, got)
	}

	dataTest2 := []int32{0, 0, 0, 0, 1, 0}
	want = 3

	got = JumpingOnCloud(dataTest2)
	if got != want {
		t.Errorf("want %v and got %v", want, got)
	}
}
