package main

import (
	"fmt"

	depone "github.com/jimmyjames85/depOne"
)

func main() {
	fmt.Printf("calling depone.One(): %s\n", depone.One())
	fmt.Printf("calling depone.OneInt(): %d\n", depone.OneInt())
}
