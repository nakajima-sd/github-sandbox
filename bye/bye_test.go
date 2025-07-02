package bye

import "testing"

func TestBye(t *testing.T) {
	want := true
	got := Bye()
	if got != want {
		t.Errorf("want %t, but got %t", want, got)
	}
}

