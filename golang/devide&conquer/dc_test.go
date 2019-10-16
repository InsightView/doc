package main

import "testing"

func Testdc(t *testing.T) {
	got := searchMatric()

	expect = true

	if got != expect{
		t.Errorf("got %b wnat %b ", got, want)
	}
}
