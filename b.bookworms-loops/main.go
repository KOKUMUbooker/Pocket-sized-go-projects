package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main(){
	// TODO: Figure out why file option from CLI doesn't override testdata/bookworms.json
	var file string;
	flag.StringVar(&file,"file","testdata/bookworms.json","Specify file path of your test data eg testdata/bookworms.json");
	flag.Parse();
	fmt.Println("file : ",file);

	bookworms,err := LoadBookworms(file);
	if err != nil {
		_,_ = fmt.Fprintf(os.Stderr,"Failed to load bookworms : %q ",err)
		os.Exit(1);
	}

	indentedJson,err := json.MarshalIndent(bookworms,""," ");
	if err != nil {
		fmt.Fprint(os.Stderr,"Failed to indent json data");
		os.Exit(1);
	}

	fmt.Println("\nBookworms : \n",string(indentedJson),"\n ");

	commonBooks := findCommonBooks(bookworms);
	fmt.Println("Here are the books in common:");
	displayBooks(commonBooks)
}