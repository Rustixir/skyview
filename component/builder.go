package component

import (
	"bytes"
	"context"
	"github.com/Rustixir/skyview/render"
	"github.com/Rustixir/skyview/server"
	"github.com/jfyne/live"
	"io"
	"strings"
)

type Builder struct {
	secret      string
	sessionName string
}

func NewBuilder(secret string, sessionName string) Builder {
	return Builder{
		secret,
		sessionName,
	}
}

// name used for url: baseurl/{lowercase name}
func (b *Builder) AddComponent(name string, factroy func() Component, listenOnEvents []string) {
	h := live.NewHandler()
	h.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		c := factroy()
		return c, c.Init()
	})
	h.HandleUnmount(func(s live.Socket) error {
		return s.Assigns().(Component).Terminate()
	})
	h.HandleRender(func(ctx context.Context, rc *live.RenderContext) (io.Reader, error) {
		var (
			err error
			buf bytes.Buffer
		)

		comp := rc.Assigns.(Component)
		name, html := comp.Render()
		if html != "" {
			err = render.Render(name, html, rc, &buf)
			return &buf, err
		}

		err = render.RenderFile(name, rc, &buf)
		return &buf, err
	})
	h.HandleSelf("broadcast", func(ctx context.Context, s live.Socket, data interface{}) (interface{}, error) {
		comp := s.Assigns().(Component)
		err := comp.HandleBroadcast(data)
		return comp, err
	})

	for _, evt := range listenOnEvents {
		h.HandleEvent(evt, apply(evt))
	}

	handler := live.NewHttpHandler(live.NewCookieStore(b.sessionName, []byte(b.secret)), h)
	server.AddHandler("/"+strings.ToLower(name), handler)
}

func (b *Builder) Start(port string) {
	server.Start(port)
}

func apply(evt string) live.EventHandler {
	return func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		comp := s.Assigns().(Component)
		resp := comp.HandleEvent(evt, p)
		if resp.BroadcastData != nil {
			s.Broadcast("broadcast", resp.BroadcastData)
		}

		return comp, resp.Err
	}
}
