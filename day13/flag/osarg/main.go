package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println(os.Args)
	fmt.Println("Hello", os.Args[1:])
}
