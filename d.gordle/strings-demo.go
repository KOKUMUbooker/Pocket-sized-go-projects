package main

import "fmt"

// Demo test for how strings can be converted to an array of runes or bytes
// and how their length & representation changes
func StringsDemo() {
	fmt.Println();

	msg := "Hello, 世界";
	fmt.Println(msg)
	fmt.Println("Msg len : ",len(msg))
	
	runes := []rune(msg);
	fmt.Println("\n[]rune msg len : ",len(runes))
	for _,r := range runes {
		fmt.Printf("%d ",r);
	}

	bytes := []byte(msg);
	fmt.Println("\n\n[]byte msg len : ",len(bytes))
	for _,b := range bytes {
		fmt.Printf("%d ",b);
	}
	fmt.Println();
}