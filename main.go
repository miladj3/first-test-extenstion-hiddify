package main

import (
	_ "github.com/miladj3/first-test-extenstion-hiddify/hiddify_extension"

	"github.com/hiddify/hiddify-core/extension/server"
)

func main() {
	server.StartTestExtensionServer()
}
