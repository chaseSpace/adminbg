package main

import (
	"flag"
	"fmt"
	"time"
)

/*
Non-business logic, Auxiliary tools
*/

func main() {
	p := flag.String("p", "", "print input")
	t := flag.Bool("time", false, "with time")

	flag.Parse()

	if *t {
		fmt.Printf("%s ", time.Now().String()[:19])
	}

	if *p != "" {
		fmt.Println("PRINT: ", *p)
	}
}
