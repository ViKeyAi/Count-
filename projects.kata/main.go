package main

import "fmt"

type Counter struct {
	N int
}

func (c *Counter) Inc() {
	c.N++
}
func main() {
	count := Counter{N: 3}
	fmt.Println(count)
	count.Inc()
	fmt.Println(count)
}
