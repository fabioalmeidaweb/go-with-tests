// Writes an SVG clockface of the current time to Stdout.
package main

import (
	"os"
	"time"

	"go-with-tests/math/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
