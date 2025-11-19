package main

import "testing"

func TestGreet_English(t *testing.T){
	lang := language("en")
	want := "Hello world";

	got := greet(lang);

	if got != want {
		t.Errorf("Expected: %q, got: %q",want,got);
	}
}

func TestGreet_French(t *testing.T){
	lang := language("fr")
	want := "Bonjour le monde";

	got := greet(lang);

	if got != want {
		t.Errorf("Expected: %q, got: %q",want,got);
	}
}

func TestGreet_Swahili(t *testing.T){
	lang := language("sw")
	want := "";

	got := greet(lang);

	if got != want {
		t.Errorf("Expected: %q, got: %q",want,got);
	}
}