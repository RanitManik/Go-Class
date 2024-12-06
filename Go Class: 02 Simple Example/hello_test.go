package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, test!"
	got := Say("test")

	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
