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

func (b *Builder) AddComponent(comp Component, listenOnEvents []string) {
	h := live.NewHandler()
	h.HandleMount(func(ctx context.Context, c live.Socket) (interface{}, error) {
		return comp, comp.Init()
	})
	h.HandleUnmount(func(c live.Socket) error {
		return comp.Terminate()
	})
	h.HandleRender(func(ctx context.Context, rc *live.RenderContext) (io.Reader, error) {
		var (
			err error
			buf bytes.Buffer
		)

		name, html := comp.Render()
		if html != "" {
			err = render.Render(name, html, rc, &buf)
			return &buf, err
		}

		err = render.RenderFile(name, rc, &buf)
		return &buf, err
	})

	for _, evt := range listenOnEvents {
		h.HandleEvent(evt, apply(comp, evt))
	}

	handler := live.NewHttpHandler(live.NewCookieStore(b.sessionName, []byte(b.secret)), h)
	server.AddHandler("/"+strings.ToLower(comp.ComponentName()), handler)
}

func apply(comp Component, evt string) live.EventHandler {
	return func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		return comp, comp.HandleEvent(evt, p)
	}
}
