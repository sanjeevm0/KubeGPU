package main

import (
	"fmt"
	"plugin"

	"github.com/Microsoft/KubeGPU/crishim/pkg/app"
)

func main() {
	// Add devices here
	if err := device.DeviceManager.CreateAndAddDevice("nvidiagpu"); err != nil {
		app.Die(fmt.Errorf("Adding device nvidiagpu fails with error %v", err))
	}
	// start the device manager
	device.DeviceManager.Start()
	// run the app - parses all command line arguments
	app.RunApp()
}
