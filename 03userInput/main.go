package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	var reader = bufio.NewReader(os.Stdin)
	test ,err := reader.ReadString('\n')
	fmt.Print(test)
	fmt.Print(err)
}