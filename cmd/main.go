package main

import (
	"fmt"

	depone "github.com/jimmyjames85/depOne"
	deptwo "github.com/jimmyjames85/depTwo/pkg"
)

func main() {
	fmt.Printf("calling depone.One(): %s\n", depone.One())
	fmt.Printf("calling depone.OneInt(): %d\n", depone.OneInt())
	fmt.Printf("calling deptwo.Two(): %s\n", deptwo.Two())
}
