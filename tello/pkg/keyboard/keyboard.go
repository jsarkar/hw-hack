package keyboard

import (
	"fmt"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/keyboard"
)

func keyBoardInput(drone *tello.Driver) {
	keys := keyboard.NewDriver()

	keys.On(keyboard.Key, func(data interface{}) {
		key := data.(keyboard.KeyEvent)
		switch key.Key{
		case keyboard.P:
			fmt.Println(key.Char)
			drone.TakeOff()
		case keyboard.Q:
			fmt.Println(key.Char)
			drone.Land()
		}

	})
}
