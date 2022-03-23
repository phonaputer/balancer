# Balancer

## Overview

This package contains a round-robin load balancer which can be used to balance load over any Go type.
This was implemented to test the new Go 1.18 generics feature!

## How to use this package

This package is very simple to use. 
Here is an example use case using a Round-robin balancer over a slice of `string` basic type elements.

```go
package main

import (
	"fmt"
	"github.com/phonaputer/balancer"
)

func main() {
	elements := []string{"a", "b", "c"}
	
	roundRobin := balancer.NewRoundRobin(elements)
	
	fmt.Println(roundRobin.Next()) // prints "a"
	fmt.Println(roundRobin.Next()) // prints "b"
	fmt.Println(roundRobin.Next()) // prints "c"
	fmt.Println(roundRobin.Next()) // prints "a"
	fmt.Println(roundRobin.Next()) // prints "b"
	fmt.Println(roundRobin.Next()) // prints "c"
	// etc...
}
```

In addition, the `Elements` function can be used to get back the original slice passed to `NewRoundRobin`. 

Please read [the Godoc](https://pkg.go.dev/github.com/phonaputer/balancer) for this package for more details.

*Copyright 2022 Phonaputer*
