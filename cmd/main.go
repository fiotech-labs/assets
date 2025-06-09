package main

import (
	"github.com/fiotech-labs/assets/internal/manager"
)

func main() {
	manager.InitCommands()
	manager.Execute()
}
