package main

import "fmt"

func DeferOrderExec(){
	defer fmt.Println("or not?");
	defer fmt.Println("a bookworm");
	fmt.Println("you are")
}