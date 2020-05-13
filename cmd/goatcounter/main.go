package main

import (
	"zgo.at/goatcounter"
	"zgo.at/goatcounter/handlers"
)

func main() {
	_, _ = goatcounter.Ref, handlers.Ref
}
