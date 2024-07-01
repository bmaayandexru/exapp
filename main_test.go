package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello!!!"
	got := Hello()
	if got != want {
		t.Errorf("want %s got %s", want, got)
	}
}

func TestMain(m *testing.M) {
	main()
}
