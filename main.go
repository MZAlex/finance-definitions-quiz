package main

import (
	c "finance-definitions-quiz/core"
	i "finance-definitions-quiz/introduction"
)

// Each step is following the previous one in a specific order.
// The author knows that it is possible to reduce the lines of code by embedding function calls to each other,
// But for the sake of clarty it is let this way
func main() {
	language := i.Introduction()
	c.Core(language)
}
