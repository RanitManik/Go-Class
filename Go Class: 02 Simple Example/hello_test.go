package hello

import "testing"

//func TestHello(t *testing.T) {
//	want := "Hello, test!"
//	got := Say("test")
//
//	if want != got {
//		t.Errorf("want %s, got %s", want, got)
//	}
//}

func TestHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, World!",
		},
		{
			items:  []string{"Ranit"},
			result: "Hello, Ranit!",
		},
		{
			items:  []string{"Ranit", "Soumen"},
			result: "Hello, Ranit, Soumen!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s", st.result, st.items, s)
		}
	}
}
