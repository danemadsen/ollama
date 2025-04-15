package main

// Compile with the following to get rid of the cmd pop up on windows
// go build -ldflags="-H windowsgui" .

import (
	"github.com/ollama/ollama/app/lifecycle"
	"github.com/ollama/ollama/server/zeroconf"
)

func main() {
	zeroconf.RegisterService()
	lifecycle.Run()
	zeroconf.UnregisterService()
}
