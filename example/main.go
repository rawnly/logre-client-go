package main

import "github.com/rawnly/go-logre"

func main() {
	logre.Init("<box_id>")
	_, _ = logre.Message("Hello world from GO!")
}
