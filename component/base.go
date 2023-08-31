package component

import (
	"github.com/jfyne/live"
)

type Component interface {
	// set to router baseUrl/{componentName}
	ComponentName() string

	// used for preparing model data for mounting
	Init() error

	// handle all event listenOn
	HandleEvent(evtName string, p live.Params) error

	// render html or if html is "" then try to find templateFile from FS by name
	Render() (name string, html string)

	// when socket closed called terminate
	Terminate() error
}
