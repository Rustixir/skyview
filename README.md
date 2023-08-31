# skyview
SkyView enables rich, real-time user experiences with server-rendered HTML.


Example: 

``` 

package main

import (
	"errors"
	"github.com/Rustixir/skyview/component"
	"github.com/Rustixir/skyview/server"
	"github.com/jfyne/live"
	"log"
)

func main() {
	builder := component.NewBuilder("secret-weak", "skyview")
	builder.AddComponent(NewThermoModel(), []string{"temp-up", "temp-down"})
	server.Start(":7072")
}

type Thermostat struct {
	C int
}

func NewThermoModel() *Thermostat {
	return &Thermostat{
		C: 1,
	}
}

func (c *Thermostat) ComponentName() string {
	return "Thermostat"
}

func (c *Thermostat) Init() error {
	log.Println("Initialize")
	c.C = 1
	return nil
}

func (c *Thermostat) HandleEvent(evtName string, p live.Params) error {
	switch evtName {
	case "temp-up":
		c.C += 1
	case "temp-down":
		c.C -= 1
	default:
		return errors.New("unknown command")
	}

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


```
