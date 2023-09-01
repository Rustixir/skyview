package main

import (
	"errors"
	"github.com/Rustixir/skyview"
	"github.com/Rustixir/skyview/component"
	"github.com/jfyne/live"
	"log"
)

func main() {
	builder := component.NewBuilder("secret-weak", "skyview")
	builder.AddComponent("Thermostat", NewThermoModel(), []string{"temp-up", "temp-down"})
	builder.Start(":7072")
}

type Thermostat struct {
	C int
}

func NewThermoModel() *Thermostat {
	return &Thermostat{
		C: 1,
	}
}

func (c *Thermostat) Init() error {
	log.Println("Initialize")
	c.C = 1
	return nil
}

func (c *Thermostat) HandleEvent(evtName string, p live.Params) skyview.Response {
	switch evtName {
	case "temp-up":
		c.C += 1
	case "temp-down":
		c.C -= 1
	default:
		err := errors.New("unknown command")
		return skyview.RaiseError(err)
	}

	return skyview.Ok()
}

func (c *Thermostat) HandleBroadcast(data interface{}) error {
	return nil
}

func (c *Thermostat) Render() (name string, html string) {

	// Render Html
	return "template_counter", `
		 <div>{{.Assigns.C}}</div>
	       <button live-click="temp-up">+</button>
	       <button live-click="temp-down">-</button>
	       <!-- Include to make live work -->
	       <script src="/live.js"></script>
	`

	// Render Html From File
	//return "../view/termostat.html", ""
}

func (c *Thermostat) Terminate() error {
	log.Println("Terminate")
	return nil
}
