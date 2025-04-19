package main

import (
	"os"
	"time"

	clockface "github.com/skillsworld46/go-tdd-sandbox/math"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
