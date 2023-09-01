package chat

import (
	"errors"
	"github.com/Rustixir/skyview"
	"github.com/Rustixir/skyview/component"
	"github.com/jfyne/live"
	"strconv"
)

type State struct {
	Username string
}

type Channel struct {
	Hub      *Hub
	Username string
}

func NewChannel(h *Hub) func() component.Component {
	return func() component.Component {
		return &Channel{
			Hub: h,
		}
	}
}

func (c *Channel) Init() error {
	username := "anonymous_" + strconv.FormatUint(c.Hub.NextID(), 10)
	c.Username = username
	c.Hub.Subscribe(username)
	return nil
}

func (c *Channel) HandleEvent(evtName string, p live.Params) skyview.Response {
	switch evtName {
	case "new_message":
		if msg := p.String("message"); len(msg) > 0 {
			message := c.Hub.NewMessage(c.Username, msg)
			return skyview.Broadcast(message)
		}
	default:
		err := errors.New("unknown command")
		return skyview.RaiseError(err)
	}

	return skyview.Ok()
}

func (c *Channel) HandleBroadcast(msg interface{}) error {
	return nil
}

func (c *Channel) Render() (name string, html string) {
	return "../layout.html", ""
}

func (c *Channel) Terminate() error {
	delete(c.Hub.Presence, c.Username)
	return nil
}
