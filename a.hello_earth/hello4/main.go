package main

import "fmt"

func main(){
	greeting := greet("en");
	fmt.Println(greeting);
}

type language string;

var phrasebook = map[language]string {
	"el": "Χαίρετε Κόσμε",
	// Greek
	"en": "Hello world",
	// English
	"fr": "Bonjour le monde", // French
	"he": "םלוע םולש",
	// Hebrew
	"ur": "ﻮﻠﯿﮨ ",
	// Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

// greet says hello to the world in various languages
func greet(l language) string {
	greeting, ok := phrasebook[l];
	if !ok {
		return fmt.Sprintf("Unsupported language: %q", l)
	}

	return  greeting;
}

