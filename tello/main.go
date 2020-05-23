package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"

	"github.com/jsarkar/hw-hack/tello/pkg/keyboard"
)

func main() {
	// Init Gobot drivers
	keys := keyboard.NewDriver()
	drone := tello.NewDriver("8890")

	work := func() {
		// Handle keyboard inputs
		keys.On(keyboard.Key, keyboard.KeyBoardInput(drone))

	}

	robot := gobot.NewRobot(
		"tello",
		[]gobot.Connection{},
		[]gobot.Device{keys, drone},
		work,
	)

	robot.Start()
}