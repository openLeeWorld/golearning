package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() { // pointer recevier of method
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string { // value receiver of method
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func main() {
	var c Counter
	fmt.Println(c.String())
	c.Increment() // (&c).Increment()
	fmt.Println(c.String())
}
