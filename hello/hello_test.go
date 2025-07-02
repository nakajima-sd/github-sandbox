package hello

import "testing"

func TestHello(t *testing.T) {
	want := true
	got := Hello()
	if got != want {
		t.Errorf("want %t, but got %t", want, got)
	}
}

