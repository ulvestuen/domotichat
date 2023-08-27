package domotics

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func ControlLight(command string) (string, error) {
	switch command {
	case "garden_light_on":
		return controlGardenLight(on)
	case "garden_light_off":
		return controlGardenLight(off)
	}
	return "Unknown command for controlling light", errors.New("unknown command for controlling lights")
}

func controlGardenLight(t Toggle) (string, error) {
	response, err := http.Get(
		fmt.Sprintf("https://maker.ifttt.com/trigger/garden_light_%s/with/key/%s",
			t,
			os.Getenv("SOMETHING")))
	if err != nil || response.StatusCode != 200 {
		return fmt.Sprintf("Something wrong happened when trying to turn the garden light %v", t), err
	}
	return fmt.Sprintf("The garden light has been turned on %v", t), nil
}

type Toggle int

const (
	off Toggle = iota
	on
)

func (t Toggle) String() string {
	switch t {
	case off:
		return "off"
	case on:
		return "on"
	default:
		return "off"
	}
}
