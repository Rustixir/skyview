
![DarkBird](https://github.com/Rustixir/darkbird/blob/main/skyview.jpeg)

# skyview
SkyView enables rich, real-time user experiences with server-rendered HTML.

## Feature highlights

SkyView brings a unified experience to building web applications. You no longer
have to split work between client and server, across different toolings, layers, and
abstractions. Instead, SkyView enriches the server with a declarative and powerful
model while keeping your code closer to your data (and ultimately your source of truth):

  * **Declarative rendering:** Render HTML on the server over WebSockets with a declarative model.

  * **Rich templating language:** Enjoy golang builtin template engine .

  * **Small payloads:** SkyView is smart enough to track changes so it only sends what the client needs, making SkyView payloads much smaller than server-rendered HTML.

  * **File uploads:** Real-time file uploads with progress indicators and image previews. Process your uploads on the fly or submit them to your desired cloud service.

  * **Loose coupling:** Reuse more code via stateful components with loosely-coupled templates, state, and event handling — a must for enterprise application development.

  * **Live navigation:** Enriched links and redirects are just more ways SkyView keeps your app light and performant. Clients load the minimum amount of content needed as users navigate around your app without any compromise in user experience.


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
