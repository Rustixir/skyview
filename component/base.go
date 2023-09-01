package component

import (
	"github.com/Rustixir/skyview"
	"github.com/jfyne/live"
)

type Factory func() Component

type Component interface {

	// used for preparing model data for mounting
	Init() error

	// handle all event listenOn
	HandleEvent(evtName string, p live.Params) skyview.Response

	// handle broadcast event
	HandleBroadcast(data interface{}) error

	// render html or if html is "" then try to find templateFile from FS by name
	Render() (name string, html string)

	// when socket closed called terminate
	Terminate() error
}
