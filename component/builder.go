package component

import (
	"bytes"
	"context"
	"github.com/Rustixir/skyview"
	"github.com/Rustixir/skyview/notify"
	"github.com/Rustixir/skyview/render"
	"github.com/Rustixir/skyview/server"
	"github.com/jfyne/live"
	"io"
	"strings"
)

type Builder struct {
	secret      string
	sessionName string
	notify      *notify.Local
}

func NewBuilder(secret string, sessionName string, notifyBuff int) Builder {
	return Builder{
		secret,
		sessionName,
		notify.NewLocal(notifyBuff),
	}
}

// name used for url: baseurl/{lowercase name}
// bidi when be true create a listener to listen on server event channel
func (b *Builder) AddComponent(name string, bidi bool, factory func() Component, listenOnEvents []string) {
	h := live.NewHandler()
	h.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		c := factory()
		if bidi {
			b.notify.StartListener(ctx, s)
		}
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
	h.HandleSelf(skyview.BroadcastEvent, func(ctx context.Context, s live.Socket, data interface{}) (interface{}, error) {
		comp := s.Assigns().(Component)
		err := comp.HandleBroadcast(data)
		return comp, err
	})
	h.HandleSelf(skyview.NotifyEvent, func(ctx context.Context, s live.Socket, data interface{}) (interface{}, error) {
		var err error
		comp := s.Assigns().(Component)
		if params, ok := data.(live.Params); ok {
			resp := comp.HandleEvent(skyview.NotifyEvent, params)
			if resp.BroadcastData != nil {
				s.Broadcast(skyview.BroadcastEvent, resp.BroadcastData)
			}
			err = resp.Err
		}
		return comp, err
	})
	for _, evt := range listenOnEvents {
		h.HandleEvent(evt, apply(evt))
	}

	handler := live.NewHttpHandler(live.NewCookieStore(b.sessionName, []byte(b.secret)), h)
	server.AddHandler("/"+strings.ToLower(name), handler)
}

func (b *Builder) GetNotify() *notify.Local {
	return b.notify
}

func (b *Builder) Start(port string) {
	server.Start(port)
}

func apply(evt string) live.EventHandler {
	return func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		comp := s.Assigns().(Component)
		resp := comp.HandleEvent(evt, p)
		if resp.BroadcastData != nil {
			s.Broadcast(skyview.BroadcastEvent, resp.BroadcastData)
		}
		return comp, resp.Err
	}
}
