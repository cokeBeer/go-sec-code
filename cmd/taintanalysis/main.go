package main

import "github.com/cokeBeer/goot/pkg/example/taint"

func main() {
	runner := taint.NewRunner("go-sec-code")
	runner.SrcPath = ""
	runner.DstPath = "go-sec-code.json"
	runner.Debug = true
	runner.Run()
}
