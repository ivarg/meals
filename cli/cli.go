package main

import "flag"

var (
	sim bool
)

func init() {
	flag.BoolVar(&sim, "sim", false, "")
	flag.Parse()
}

func main() {
	meals.Generate(5)
}
