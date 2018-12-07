package main

import (
	"fmt"

	depone "github.com/jimmyjames85/depOne"
	deptwo "github.com/jimmyjames85/depTwo"
)

func main() {
	fmt.Printf("calling depone.One(): %s\n", depone.One())
	fmt.Printf("calling deptwo.Two(): %s\n", deptwo.Two())
}
